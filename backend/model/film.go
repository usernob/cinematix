package model

import (
	// "fmt"
	"fmt"
	"time"
)

type Film struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	Title        string        `json:"title"`
	Rating       float32       `json:"rating"`
	TanggalRilis time.Time     `json:"tanggal_rilis"`
	Sinopsis     string        `json:"sinopsis"`
	PosterPath   string        `json:"poster_path"`
	Genre        []*Genre      `gorm:"many2many:film_genres;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"genre,omitempty"`
	Penayangan   []*Penayangan `json:"penayangan,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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

func GetFilmTayang(cooming_soon bool) ([]Film, error) {
	var operator string
	if cooming_soon {
		operator = ">"
	} else {
		operator = "<"
	}

	var films []Film

	res := Db.InnerJoins(
		fmt.Sprintf("inner join penayangan on penayangan.film_id = films.id and penayangan.mulai %s ?", operator),
		time.Now().Add(time.Hour*24*7),
	).
		Group("films.id").
		Find(&films)

	return films, res.Error
}

func GetFilmHaveTayang() ([]Film, error) {
	var films []Film

	res := Db.InnerJoins("inner join penayangan on penayangan.film_id = films.id").Group("films.id").Find(&films)

	return films, res.Error
}

func GetFilmDetail(id string) (Film, error) {
	var film Film
	res := Db.Preload("Genre").Preload("Penayangan").Where("id = ?", id).First(&film)
	return film, res.Error
}

func GetPenayanganSingle(id string, penayangan_id string) (Film, error) {
	var film Film
	res := Db.Preload("Genre").Preload("Penayangan", "id = ?", penayangan_id).Where("id = ?", id).First(&film)
	return film, res.Error
}
