package model

import (
	"backend/pkg/logjson"
	"time"
)

type Penayangan struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	FilmID        uint           `json:"film_id"`
	Film          *Film          `json:"film,omitempty"`
	AudiotoriumID uint           `json:"audiotorium_id"`
	Harga         uint           `json:"harga"`
	Mulai         time.Time      `json:"mulai"`
	Selesai       time.Time      `json:"selesai"`
	Tiket         []*Tiket       `json:"tiket,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (p *Penayangan) TableName() string {
	return "penayangan"
}

func GetPenayangan(penayangan *Penayangan) (*Penayangan, error) {
	res := Db.Find(penayangan)
	if res.Error != nil {
		return nil, res.Error
	}
	return penayangan, nil
}

func GetKursi(penayanganId uint) ([]*Kursi, error) {
	var kursi []*Kursi
	penayangan := &Penayangan{ID: penayanganId}
	penayangan, err := GetPenayangan(penayangan)
  logjson.ToJSON(penayangan)
	if err != nil {
		return nil, err
	}

  var tiket []Tiket
  Db.Where("penayangan_id = ?", penayangan.ID).Find(&tiket)
  logjson.ToJSON(tiket)
  res := Db.Preload("Tiket", Db.Where("penayangan_id = ?", penayangan.ID)).Where("audiotorium_id = ?", penayangan.AudiotoriumID).Find(&kursi)

  logjson.ToJSON(kursi)
	if res.Error != nil {
		return nil, res.Error
	}
	return kursi, nil
}
