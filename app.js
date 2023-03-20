const form = document.querySelector('#search-form');
const animeList = document.querySelector('#anime-list');

form.addEventListener('submit', async (e) => {
	e.preventDefault();

	const query = e.target.elements[0].value;
	if (!query) return;

	const url = `https://api.jikan.moe/v3/search/anime?q=${query}&page=1`;
	const response = await fetch(url);
	const data = await response.json();
	const results = data.results;

	animeList.innerHTML = '';

    console.log(results);

	results.forEach((anime) => {
		const li = document.createElement('li');
		const img = document.createElement('img');
		const h2 = document.createElement('h2');
		const p = document.createElement('p');

		img.src = anime.image_url;
		h2.textContent = anime.title;
		p.textContent = anime.synopsis;

		li.appendChild(img);
		li.appendChild(h2);
		li.appendChild(p);

		animeList.appendChild(li);
	});

   
});
