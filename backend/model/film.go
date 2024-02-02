package model

import (
	"time"

	"gorm.io/gorm"
)

type Film struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Title        string         `json:"title"`
	Rating       float32        `json:"rating"`
	TanggalRilis time.Time      `json:"tanggal_rilis"`
	Sinopsis     string         `json:"sinopsis"`
	PosterPath   string         `json:"poster_path"`
	Genre        []*Genre       `gorm:"many2many:film_genres" json:"genre,omitempty"`
	Penayangan   []*Penayangan  `json:"penayangan,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Genre struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Nama      string         `gorm:"unique;not null;index" json:"nama"`
	Film      []*Film        `gorm:"many2many:film_genres" json:",omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type FilmGenre struct {
  FilmID    uint           `gorm:"primaryKey" json:"film_id"`
  GenreID   uint           `gorm:"primaryKey" json:"genre_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
