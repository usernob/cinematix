package model

import (
	"time"

	"gorm.io/gorm"
)

type Kursi struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Nama          string         `json:"nama" gorm:"unique;not null;index"`
	AudiotoriumID uint           `json:"audiotorium_id"`
	Tiket         []*Tiket       `gorm:"many2many:seats" json:"tiket,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Audiotorium struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Nama       string         `json:"nama" gorm:"unique;not null;index"`
	Kursi      []*Kursi       `json:"kursi,omitempty"`
	Penayangan []*Penayangan  `json:"penayangan,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (k *Kursi) TableName() string {
	return "kursi"
}

func (a *Audiotorium) TableName() string {
	return "audiotorium"
}
