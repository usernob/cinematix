package model

import "gorm.io/gorm"


type Kursi struct {
  gorm.Model
  Nama string
  AudiotoriumID uint
  Seat []*Seat
}

type Audiotorium struct {
  gorm.Model
  Nama string
  Kursi []*Kursi `gorm:"foreignKey:AudiotoriumID"`
  Penayangan []*Penayangan `gorm:"foreignKey:AudiotoriumID"`
}
