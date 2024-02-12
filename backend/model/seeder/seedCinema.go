package seeder

import (
	"backend/model"
	"fmt"
)

func seedCinema() {
	var kursis []model.Kursi

	var abjd string = "ABCDEFGHI"

	for j := 1; j < 4; j++ {
		for _, char := range abjd {
			for i := 1; i < 10; i++ {
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
	for i := 0; i < 6; i++ {
		//
		var kursi []*model.Kursi
		var penayangan model.Penayangan
		model.Db.Find(&penayangan, i+1)
		model.Db.Where("audiotorium_id = ?", penayangan.AudiotoriumID).Find(&kursi)

    tiket := model.Tiket{TotalHarga: 100000, UserID: 1, PenayanganID: penayangan.ID, StatusPembayaran: model.Waiting}
    model.Db.Create(&tiket)
    model.Db.Model(&tiket).Association("Kursi").Append(kursi[2:6])
	}
}
