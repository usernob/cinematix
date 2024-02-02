package model

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type Tiket struct {
	ID               uint              `gorm:"primaryKey" json:"id"`
	TotalHarga       uint              `json:"total_harga"`
	StatusPembayaran Status_Pembayaran `sql:"type:status_pembayaran" json:"status_pembayaran"`
	UserID           uint              `json:"user_id"`
	PenayanganID     uint              `json:"penyangan_id"`
	Seat             []*Seat           `json:"seat,omitempty"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	DeletedAt        gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
}

type Status_Pembayaran string

func (p *Status_Pembayaran) Scan(value interface{}) error {
	*p = Status_Pembayaran(value.([]byte))
	return nil
}

func (p Status_Pembayaran) Value() (driver.Value, error) {
	return string(p), nil
}

const (
	Waiting Status_Pembayaran = "waiting"
	Done    Status_Pembayaran = "done"
)
