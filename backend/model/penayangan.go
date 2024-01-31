package model

import (
	"time"

	"gorm.io/gorm"
)

type Penayangan struct {
  gorm.Model
  FilmID uint
  AudiotoriumID uint
  Harga uint
  Mulai time.Time
  Selesai time.Time
  Tiket []*Tiket
}

type Seat struct {
  gorm.Model
  KursiID uint
  TiketID uint
}
