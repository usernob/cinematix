package model

import (
	"time"

	"gorm.io/gorm"
)

type Film struct {
  gorm.Model
  Title string
  Rating float64
  TanggalRilis time.Time
  Sinopsis string
  Poster string
  Genre []*Genre `gorm:"many2many:film_genres"`
  Penayangan []*Penayangan
}

type Genre struct {
  Id uint `gorm:"primaryKey"`
  Nama string `gorm:"unique;not null;index"`
  Film []*Film `gorm:"many2many:film_genres"`
}

