package seeder

import (
	"backend/model"
	"encoding/json"
	"time"
)

func seedFilmAndPenyangan() {
	type Films struct {
		Results []struct {
			ID          uint    `json:"id"`
			Title       string  `json:"title"`
			PosterPath  string  `json:"poster_path"`
			TangalRilis string  `json:"release_date"`
			Rating      float32 `json:"vote_average"`
			Sinopsis    string  `json:"overview"`
			GenreId     []int   `json:"genre_ids"`
		} `json:"results"`
	}

	bodyString, err := getDataFromApi("/movie/now_playing?language=en-US&page=1")
	if err != nil {
		panic(err)
	}

	var resFilms Films
	json.Unmarshal([]byte(bodyString), &resFilms)

	var films []model.Film
	for _, film := range resFilms.Results {
		// randomInt := rand.Intn(3)
		tanggalRilis, err := time.Parse("2006-01-02", film.TangalRilis)
		if err != nil {
			panic(err)
		}

		// penyangan := []*model.Penayangan{}
		// if i%randomInt == 0 {
		// 	penyangan = append(penyangan, &model.Penayangan{
		// 		AudiotoriumID: uint(randomInt % 3),
		// 		Harga:         50000,
		// 		Mulai:         time.Now(),
		// 		Selesai:       time.Now().Add(time.Hour * 1),
		// 	})
		// }

		genres := []*model.Genre{}
		for _, genreId := range film.GenreId {
			genre := model.Genre{ID: uint(genreId)}
			genres = append(genres, &genre)
		}

		films = append(films, model.Film{
			ID:           film.ID,
			Title:        film.Title,
			TanggalRilis: tanggalRilis,
			Rating:       film.Rating,
			Sinopsis:     film.Sinopsis,
			PosterPath:   film.PosterPath,
			Genre:        genres,
			// Penayangan:   penyangan,
		})
	}

	model.Db.Create(films)
}
