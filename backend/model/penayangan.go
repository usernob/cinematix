package model

import (
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
	Tanggal       time.Time      `json:"tanggal"`
	Mulai         time.Time      `json:"mulai"`
	Selesai       time.Time      `json:"selesai"`
	Tiket         []*Tiket       `json:"tiket,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Report        []*Report      `json:"report,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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

func GetListPenayangan() ([]Penayangan, error) {
	var penayangans []Penayangan
	res := Db.Preload("Film").Find(&penayangans)
	return penayangans, res.Error
}

func GetPenayanganDetail(id uint) (Penayangan, error) {
	var penayangan Penayangan
	res := Db.Preload("Film").Where("id = ?", id).First(&penayangan)
	return penayangan, res.Error
}

func GetKursi(penayanganId uint) ([]*Kursi, error) {
	var kursi []*Kursi
	penayangan := &Penayangan{ID: penayanganId}
	penayangan, err := GetPenayangan(penayangan)
	if err != nil {
		return nil, err
	}

	var tiket []Tiket
	Db.Where("penayangan_id = ?", penayangan.ID).Find(&tiket)
	res := Db.Preload("Tiket.Kursi").Preload("Tiket", Db.Where("penayangan_id = ?", penayangan.ID)).Where("audiotorium_id = ?", penayangan.AudiotoriumID).Find(&kursi)

	if res.Error != nil {
		return nil, res.Error
	}
	return kursi, nil
}

func DeleteExpPenayangan() error {
	var penayangans []Penayangan

	res := Db.Preload("Tiket.Kursi").Preload("Tiket").Where("selesai < ?", time.Now()).Find(&penayangans)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	if res.Error != nil {
		return res.Error
	}

	if len(penayangans) <= 0 {
		return nil
	}

	for _, penayangan := range penayangans {
		var tiketTerjual uint
		var totalPendapatan uint
		deletePenayangan := Db.Select("Tiket").Delete(&penayangan)
		if deletePenayangan.Error != nil {
			return deletePenayangan.Error
		}

		if len(penayangan.Tiket) <= 0 {
			continue
		}

		for _, tiket := range penayangan.Tiket {
			totalPendapatan += tiket.TotalHarga
			tiketTerjual += uint(len(tiket.Kursi))
		}

		create := Db.Create(&Report{PenayanganID: penayangan.ID, TiketTerjual: tiketTerjual, Pendapatan: totalPendapatan, CreatedAt: penayangan.Selesai})
		if create.Error != nil {
			return create.Error
		}
	}

	return nil
}

func checkJadwalPenayangan(id uint, tanggal time.Time, mulai time.Time, selesai time.Time, audiotorium_id uint) bool {
	var penayangans []Penayangan
	res := Db.Where("id != ?", id).Where("audiotorium_id = ?", audiotorium_id).Where("tanggal = ?", tanggal).Where("mulai BETWEEN ? AND ?", mulai, selesai).Or("selesai BETWEEN ? AND ?", mulai, selesai).Find(&penayangans)
  if res.Error != nil {
    return true
  }
	return len(penayangans) > 0 
}


func AddPenayangan(film_id uint, tanggal time.Time, mulai time.Time, selesai time.Time, audiotorium_id uint, harga uint) error {
  if checkJadwalPenayangan(0, tanggal, mulai, selesai, audiotorium_id) {
    return errors.New("jadwal penayangan sudah ada")
  }
	penayangan := Penayangan{
		FilmID:        film_id,
		Tanggal:       tanggal,
		Mulai:         mulai,
		Selesai:       selesai,
		AudiotoriumID: audiotorium_id,
		Harga:         harga,
	}
	return Db.Create(&penayangan).Error
}

func UpdatePenayangan(id uint, film_id uint, tanggal time.Time, mulai time.Time, selesai time.Time, audiotorium_id uint, harga uint) error {
  if checkJadwalPenayangan(id, tanggal, mulai, selesai, audiotorium_id) {
    return errors.New("jadwal penayangan sudah ada")
  }
	penayangan := Penayangan{
		FilmID:        film_id,
		Tanggal:       tanggal,
		Mulai:         mulai,
		Selesai:       selesai,
		AudiotoriumID: audiotorium_id,
		Harga:         harga,
	}
	return Db.Model(&Penayangan{ID: id}).Updates(&penayangan).Error
}

func DeletePenayangan(id uint) error {
	penayangan := Penayangan{ID: id}
	return Db.Unscoped().Delete(&penayangan).Error
}
