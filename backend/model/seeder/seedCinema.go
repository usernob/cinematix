package seeder

import (
	"backend/model"
	"fmt"
	"math/rand"
)

func seedCinema() {
	var kursis []model.Kursi

	var abjd string = "ABCDEFGHI"

	for j := 1; j < 4; j++ {
		for _, char := range abjd {
			for i := 1; i <= 10; i++ {
				var kursi model.Kursi
				kursi.Nama = fmt.Sprintf("A%d%c%d", j, char, i)
				kursi.AudiotoriumID = uint(j)
				kursis = append(kursis, kursi)
			}
		}
	}
	model.Db.Create(kursis)
}

func seedTiketAndSeat() {
	var penayangans []model.Penayangan
	model.Db.Find(&penayangans)
	for _, penayangan := range penayangans {
		var kursi []*model.Kursi
		model.Db.Where("audiotorium_id = ?", penayangan.AudiotoriumID).Find(&kursi)

		offset := 0
		for i := 0; i < rand.Intn(10); i++ {
			ran := 1 + rand.Intn(5)
			tiket := model.Tiket{TotalHarga: uint(100000 * ran), UserID: 1, PenayanganID: penayangan.ID, StatusPembayaran: model.Done}
			model.Db.Create(&tiket)
			model.Db.Model(&tiket).Association("Kursi").Append(kursi[offset : offset+ran])
			offset += ran
		}
	}
}
