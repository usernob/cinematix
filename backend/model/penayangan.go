package model

import (
	"time"

	"gorm.io/gorm"
)

type Penayangan struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	FilmID        uint           `json:"film_id"`
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
