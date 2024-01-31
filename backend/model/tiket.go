package model

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type Tiket struct {
  gorm.Model
  TotalHarga uint
  StatusPembayaran Status_Pembayaran `sql:"type:status_pembayaran"`
  UserID uint
  PenayanganID uint
  Seat []*Seat
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
  Done Status_Pembayaran  = "done"
)

