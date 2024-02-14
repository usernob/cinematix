package model

import (
	"time"
)

type Kursi struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Nama          string    `json:"nama" gorm:"unique;not null;index"`
	AudiotoriumID uint      `json:"audiotorium_id"`
	Tiket         []*Tiket  `gorm:"many2many:seats;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"tiket,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Audiotorium struct {
	ID         uint          `gorm:"primaryKey" json:"id"`
	Nama       string        `json:"nama" gorm:"unique;not null;index"`
	Kursi      []*Kursi      `json:"kursi,omitempty"`
	Penayangan []*Penayangan `json:"penayangan,omitempty"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
}

type Report struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	PenayanganID uint      `json:"penayangan_id"`
	TiketTerjual uint      `json:"tiket_terjual"`
	Pendapatan   uint      `json:"pendapatan"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (k *Kursi) TableName() string {
	return "kursi"
}

func (a *Audiotorium) TableName() string {
	return "audiotorium"
}
