package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Anime struct {
	Title       string `json:"title"`
	ImageURL    string `json:"image_url"`
	Type        string `json:"type"`
	Episodes    int    `json:"episodes"`
	Score       string `json:"score"`
	Description string `json:"synopsis"`
}

// Fonction pour récupérer les données de l'API MyAnimeList
func getAnimeData(search string) ([]Anime, error) {
	// Remplacer "API_KEY" par votre propre clé API MyAnimeList
	url := fmt.Sprintf("https://api.myanimelist.net/v2/anime?q=%s&limit=10", search)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Remplacer "ACCESS_TOKEN" par votre propre jeton d'accès MyAnimeList
	req.Header.Set("Authorization", "Bearer ACCESS_TOKEN")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string][]Anime
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result["data"], nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			search := r.FormValue("search")
			search = strings.ReplaceAll(search, " ", "%20")
			animeData, err := getAnimeData(search)
			if err != nil {
				log.Println(err)
				return
			}

			t, err := template.ParseFiles("index.html")
			if err != nil {
				log.Println(err)
				return
			}

			err = t.Execute(w, animeData)
			if err != nil {
				log.Println(err)
				return
			}
		} else {
			http.ServeFile(w, r, "index.html")
		}
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
