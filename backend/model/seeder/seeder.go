package seeder

import (
	"backend/model"
	"backend/pkg/filewriter"
	"backend/pkg/passwordhash"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func getDataFromApi(path string) (string, error) {

	url := os.Getenv("TMDB_API_BASE_URL") + path

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TMDB_API_KEY")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func downloadPoster(posterPath string) (string, error) {
	url := os.Getenv("TMDB_API_IMAGE_BASE_URL") + posterPath

	path, filename := filewriter.SplitPath("storage/image/poster" + posterPath)

	filePath := filepath.Join(filepath.Join(path...), filename)

	if filewriter.CheckFileExists(filePath) {
		return filePath, nil
	}

	fmt.Println("GET: ", url)
	fmt.Println("DEST: ", filePath)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil
	}
	defer res.Body.Close()

	mkdirErr := os.MkdirAll(filepath.Join(path...), 0755)
	if mkdirErr != nil {
		return "", mkdirErr
	}

	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	defer out.Close()

	_, copyErr := io.Copy(out, res.Body)
	if copyErr != nil {
		return "", copyErr
	}

	return filePath, nil
}

func Seed() {
	pass, err := passwordhash.HashPassword("password")
	if err != nil {
		panic(err)
	}

	model.Db.Create(&model.User{
		Nama:     "user1",
		Email:    "user1@gmail.com",
		Password: pass,
	})

	seedGenre()
	seedFilmAndPenyangan()
}
