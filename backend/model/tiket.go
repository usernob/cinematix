package model

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Tiket struct {
	ID               uint              `gorm:"primaryKey" json:"id"`
	TotalHarga       uint              `json:"total_harga"`
	StatusPembayaran Status_Pembayaran `sql:"type:status_pembayaran" json:"status_pembayaran"`
	UserID           uint              `json:"user_id"`
	PenayanganID     uint              `json:"penyangan_id"`
	Kursi            []*Kursi          `gorm:"many2many:seats;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"kursi,omitempty"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	DeletedAt        gorm.DeletedAt    `gorm:"index"`
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
	tiket.StatusPembayaran = Waiting
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

func UpdateTiketPembayaran(tiket *Tiket) error {
	err := Db.Model(tiket).Update("status_pembayaran", Done)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func GetPesanan(tiketid uint, userid uint, status Status_Pembayaran) (*Film, error) {
	var penayangan Penayangan
	var film Film

	err := Db.Preload("Tiket.Kursi").
		Preload("Tiket", Db.Where(&Tiket{ID: tiketid, UserID: userid, StatusPembayaran: status})).
		Where("id IN (?)", Db.Model(&Tiket{ID: tiketid, UserID: userid, StatusPembayaran: status}).Select("penayangan_id")).
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

func GetAllTiket(userid uint) ([]*Tiket, error) {
	var tikets []*Tiket
	err := Db.Where(&Tiket{UserID: userid}).Find(&tikets)
	if err.Error != nil {
		return nil, err.Error
	}
	return tikets, nil
}

func DeleteExpTiket() {
	var tikets []*Tiket
	var tiketExp []Tiket
	res := Db.Where(&Tiket{StatusPembayaran: Waiting}).Find(&tikets)
	if res.Error != nil {
		return
	}

	for _, tik := range tikets {
		if tik.UpdatedAt.Add(time.Minute * 2).Before(time.Now()) {
			tiketExp = append(tiketExp, *tik)
			fmt.Println(tik.UpdatedAt, time.Now())
		}
	}
	if len(tiketExp) > 0 {
		Db.Select("Kursi").Delete(&tiketExp)
	}
}
