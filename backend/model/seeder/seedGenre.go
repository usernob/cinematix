package seeder

import (
	"backend/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getDataFromApi(path string) (string, error) {

	url := fmt.Sprintf("%s%s", os.Getenv("TMDB_API_BASE_URL"), path)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TMDB_API_KEY")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil
	}

	return string(body), nil
}

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
