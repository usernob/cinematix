package model

import (
	"backend/pkg/passwordhash"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Seed() {
	pass, err := passwordhash.HashPassword("password")
	if err != nil {
		panic(err)
	}

	db.Create(&User{
		Nama:     "user1",
		Email:    "user1@gmail.com",
		Password: pass,
	})

	db.Create(getGenres())
}

func getGenres() []Genre {

	type Genres struct {
		Genres []struct {
			Id   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"genres"`
	}

	url := fmt.Sprintf("%s/genre/movie/list?language=en", os.Getenv("TMDB_API_BASE_URL"))

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TMDB_API_KEY")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(body)

	var resGenres Genres
	// fmt.Println(bodyString)
	json.Unmarshal([]byte(bodyString), &resGenres)

	var genres []Genre
	for _, genre := range resGenres.Genres {
		genres = append(genres, Genre{
			Id:   genre.Id,
			Nama: genre.Name,
		})
	}

	return genres
}
