package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AnimeResult struct {
	ImageURL string `json:"image_url"`
	Title    string `json:"title"`
	Synopsis string `json:"synopsis"`
}

type AnimeSearchResults struct {
	Results []AnimeResult `json:"results"`
}

func main() {
	http.HandleFunc("/search", animeSearchHandler)
	http.ListenAndServe(":8080", nil)
}

func animeSearchHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	searchQuery := query.Get("q")
	if searchQuery == "" {
		http.Error(w, "Missing search query parameter", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://api.jikan.moe/v3/search/anime?q=%s&page=1", searchQuery)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var searchResults AnimeSearchResults
	err = json.Unmarshal(body, &searchResults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(searchResults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
