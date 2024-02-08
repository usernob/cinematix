package model

import (
	"fmt"
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
	Tiket         []*Tiket       `json:"tiket,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (p *Penayangan) TableName() string {
	return "penayangan"
}

type Seat struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	KursiID   uint           `json:"kursi_id"`
	TiketID   uint           `json:"tiket_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
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
	fmt.Println(penayangan)
	if err != nil {
		return nil, err
	}

	res := Db.Preload("Seat", "tiket_id IN (?)", Db.Model(&Tiket{}).Select("id").Where("penayangan_id = ?", penayangan.ID)).
		Where("audiotorium_id = ?", penayangan.AudiotoriumID).
		Find(&kursi)

	if res.Error != nil {
		return nil, res.Error
	}
	return kursi, nil
}
