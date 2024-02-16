package seeder

import (
	"backend/model"
	"encoding/json"
	"math/rand"
	"time"
)

func seedFilmAndPenyangan() {
	audiotorium := []model.Audiotorium{
		{
			Nama: "Audiotorium 1",
		},
		{
			Nama: "Audiotorium 2",
		},
		{
			Nama: "Audiotorium 3",
		},
	}
	model.Db.Create(audiotorium)
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
	for i, film := range resFilms.Results {
		tanggalRilis, err := time.Parse("2006-01-02", film.TangalRilis)
		if err != nil {
			panic(err)
		}

		penyangan := []*model.Penayangan{}
		for j := 0; j < rand.Intn(25); j++ {
			var mulai time.Time
			var selesai time.Time

			if j%3 == 0 {
				mulai = time.Now().Add(time.Hour * (4 * time.Duration(-i+-j)))
				selesai = time.Now().Add(time.Hour*(4*time.Duration(-i+-j)) + time.Hour*2)
			} else {
				mulai = time.Now().Add(time.Hour * (12 * time.Duration(i+j)))
				selesai = time.Now().Add(time.Hour*(12*time.Duration(i+j)) + time.Hour*2)
			}

			penyangan = append(penyangan, &model.Penayangan{
				AudiotoriumID: audiotorium[rand.Intn(3)].ID,
				Harga:         50000,
				Mulai:         mulai,
				Selesai:       selesai,
			})
		}

		genres := []*model.Genre{}
		for _, genreId := range film.GenreId {
			genre := model.Genre{ID: uint(genreId)}
			genres = append(genres, &genre)
		}

		downloadPath, err := downloadPoster(film.PosterPath)
		if err != nil {
			panic(err)
		}

		films = append(films, model.Film{
			ID:           film.ID,
			Title:        film.Title,
			TanggalRilis: tanggalRilis,
			Rating:       film.Rating,
			Popularitas:  uint(rand.Intn(100)),
			Sinopsis:     film.Sinopsis,
			PosterPath:   downloadPath,
			Genre:        genres,
			Penayangan:   penyangan,
		})
	}

	model.Db.Create(films)
}
