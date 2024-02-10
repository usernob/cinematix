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
	Kursi            []*Kursi          `gorm:"many2many:seats" json:"kursi,omitempty"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	DeletedAt        gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
}

type Seats struct {
	KursiID uint `json:"kursi_id"`
	TiketID uint `json:"tiket_id"`
}

type Status_Pembayaran string

func (p *Status_Pembayaran) Scan(value interface{}) error {
	*p = Status_Pembayaran(value.(string))
	return nil
}

func (p Status_Pembayaran) Value() (driver.Value, error) {
	return string(p), nil
}

const (
	Waiting Status_Pembayaran = "waiting"
	Done    Status_Pembayaran = "done"
)

func CreateTiketAndSeat(tiket *Tiket, seat []*Kursi) error {
	err := Db.Create(tiket)
	if err.Error != nil {
		return err.Error
	}
	appendErr := Db.Model(tiket).Association("Kursi").Append(seat)
	if appendErr != nil {
		return appendErr
	}
	return nil
}

func GetPesanan(tiketid uint, userid uint) (*Film, error) {
	var penayangan Penayangan
	var film Film

	err := Db.Preload("Tiket.Kursi").
		Preload("Tiket", Db.Where(&Tiket{ID: tiketid, UserID: userid})).
		Where("id IN (?)", Db.Model(&Tiket{ID: tiketid, UserID: userid}).Select("penayangan_id")).
		Find(&penayangan)
	if err.Error != nil {
		return nil, err.Error
	}

	Filmerr := Db.Where("id = ?", penayangan.FilmID).Find(&film)
	if Filmerr.Error != nil {
		return nil, Filmerr.Error
	}

  film.Penayangan = append(film.Penayangan, &penayangan)

	return &film, nil
}
