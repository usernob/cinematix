package model

import (
	// "fmt"

	"fmt"
	"os"
	"time"

	"gorm.io/gorm"
)

type Film struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	Title        string        `json:"title"`
	Rating       float32       `json:"rating"`
	TanggalRilis time.Time     `json:"tanggal_rilis"`
	Sinopsis     string        `json:"sinopsis"`
	PosterPath   string        `json:"poster_path"`
	Popularitas  uint          `json:"popularitas" gorm:"default:0"`
	Genre        []*Genre      `gorm:"many2many:film_genres;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"genre,omitempty"`
	Penayangan   []*Penayangan `json:"penayangan,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

type Genre struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"unique;not null;index" json:"nama"`
	Film      []*Film   `gorm:"many2many:film_genres;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:",omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FilmGenre struct {
	FilmID    uint      `gorm:"primaryKey" json:"film_id"`
	GenreID   uint      `gorm:"primaryKey" json:"genre_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResFilmTayang struct {
	Penayangan
	FilmID         uint   `json:"film_id"`
	FilmTitle      string `json:"film_title"`
	FilmPosterPath string `json:"film_poster_path"`
}

func GetAllFilm() ([]Film, error) {
	var films []Film
	res := Db.Find(&films)
	if res.Error != nil {
		return nil, res.Error
	}
	return films, nil
}

func GetFilmTayang(cooming_soon bool) ([]Film, error) {
	var operator string
	if cooming_soon {
		operator = ">"
	} else {
		operator = "<"
	}

	var films []Film

	res := Db.InnerJoins(
		fmt.Sprintf("inner join penayangan on penayangan.film_id = films.id and penayangan.mulai %s ? and penayangan.deleted_at is null", operator),
		time.Now().Add(time.Hour*24*7),
	).
		Group("films.id").
		Find(&films)

	return films, res.Error
}

func GetFilmPopuler() ([]Film, error) {
	var films []Film
	res := Db.InnerJoins("inner join penayangan on penayangan.film_id = films.id and penayangan.deleted_at is null").
		Group("films.id").
		Order("popularitas desc").
		Order("rating desc").
		Limit(10).Find(&films)
	return films, res.Error
}

func GetFilmHaveTayang() ([]Film, error) {
	var films []Film

	res := Db.InnerJoins("inner join penayangan on penayangan.film_id = films.id and penayangan.deleted_at is null").Group("films.id").Find(&films)

	return films, res.Error
}

func GetFilmDetail(id uint) (Film, error) {
	var film Film
	res := Db.Preload("Genre").Preload("Penayangan").Where("id = ?", id).First(&film)
	return film, res.Error
}

func GetPenayanganSingle(id string, penayangan_id string) (Film, error) {
	var film Film
	res := Db.Preload("Genre").Preload("Penayangan", "id = ?", penayangan_id).Where("id = ?", id).First(&film)
	return film, res.Error
}

func AddFilm(title string, rating float32, tanggalRilis time.Time, sinopsis string, posterPath string, genreIds []int) error {
	film := Film{
		Title:        title,
		Rating:       rating,
		TanggalRilis: tanggalRilis,
		Sinopsis:     sinopsis,
		PosterPath:   posterPath,
	}

	err := Db.Create(&film)
	if err.Error != nil {
		return err.Error
	}

	if len(genreIds) > 0 {
		genres := []Genre{}
		for _, genreId := range genreIds {
			genre := Genre{ID: uint(genreId)}
			genres = append(genres, genre)
		}

		err := Db.Model(&film).Association("Genre").Append(genres)
		if err != nil {
			return err
		}
	}
	return nil
}

func EditFilm(id int, title string, rating float32, tanggalRilis time.Time, sinopsis string, posterPath string, genreIds []int) error {
	fmt.Println(posterPath)
	fmt.Println(genreIds)
	film, err := GetFilmDetail(uint(id))
	if err != nil {
		return err
	}

	film.Title = title
	film.Rating = rating
	film.TanggalRilis = tanggalRilis
	film.Sinopsis = sinopsis

	if posterPath != "" {
		os.Remove(film.PosterPath)
		film.PosterPath = posterPath
	}

	uErr := Db.Model(&Film{ID: uint(id)}).Updates(&film)
	if uErr.Error != nil {
		return uErr.Error
	}

	if len(genreIds) > 0 {
		genres := []Genre{}
		for _, genreId := range genreIds {
			genre := Genre{ID: uint(genreId)}
			genres = append(genres, genre)
		}

		err := Db.Model(&film).Association("Genre").Replace(genres)
		if err != nil {
			return err
		}
	} else {
		err := Db.Model(&film).Association("Genre").Delete(&film.Genre)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteFilm(id uint) error {
	res := Db.Where("id = ?", id).Delete(&Film{})
	return res.Error
}

func FindGenres(query string) ([]Genre, error) {
	var genres []Genre
	res := Db.Where("nama like ?", "%"+query+"%").Find(&genres)
	fmt.Println(Db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("LOWER(nama) like ?", "LOWER(%"+query+"%)").Find(&genres)
	}))
	return genres, res.Error
}
