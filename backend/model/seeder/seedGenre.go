package seeder

import (
	"backend/model"
	"encoding/json"
)

func seedGenre() {
	type Genres struct {
		Genres []struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"genres"`
	}

	bodyString, err := getDataFromApi("/genre/movie/list?language=en")
	if err != nil {
		panic(err)
	}

	var resGenres Genres
	json.Unmarshal([]byte(bodyString), &resGenres)

	var genres []model.Genre
	for _, genre := range resGenres.Genres {
		genres = append(genres, model.Genre{
			ID:   genre.Id,
			Nama: genre.Name,
		})
	}

	model.Db.Create(genres)
}
