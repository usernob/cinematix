package model

import (
	"backend/pkg/logjson"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Penayangan struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	FilmID        uint           `json:"film_id"`
	Film          *Film          `json:"film,omitempty"`
	AudiotoriumID uint           `json:"audiotorium_id"`
	Harga         uint           `json:"harga"`
	Mulai         time.Time      `json:"mulai"`
	Selesai       time.Time      `json:"selesai"`
	Tiket         []*Tiket       `json:"tiket,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Report        []*Report      `json:"report,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
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
	res := Db.Preload("Tiket.Kursi").Preload("Tiket", Db.Where("penayangan_id = ?", penayangan.ID)).Where("audiotorium_id = ?", penayangan.AudiotoriumID).Find(&kursi)

	logjson.ToJSON(kursi)
	if res.Error != nil {
		return nil, res.Error
	}
	return kursi, nil
}

func DeleteExpPenayangan() error {
	var penayangan Penayangan
	var tiketTerjual uint
	var totalPendapatan uint

	res := Db.Preload("Tiket.Kursi").Preload("Tiket").Where("selesai < ?", time.Now()).First(&penayangan)

  if errors.Is(res.Error, gorm.ErrRecordNotFound) {
    return nil 
  }

	if res.Error != nil {
		return res.Error
	}

  logjson.ToJSON(penayangan)

	for _, tiket := range penayangan.Tiket {
		totalPendapatan += tiket.TotalHarga
		tiketTerjual += uint(len(tiket.Kursi))
	}

  deletePenayangan := Db.Select("Tiket").Delete(&penayangan)
  if deletePenayangan.Error != nil {
    return deletePenayangan.Error
  }

	create := Db.Create(&Report{PenayanganID: penayangan.ID, TiketTerjual: tiketTerjual, Pendapatan: totalPendapatan})
	if create.Error != nil {
		return create.Error
	}


	return nil
}
