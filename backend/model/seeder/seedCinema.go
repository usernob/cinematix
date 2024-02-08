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
	var tiket []model.Tiket
	for i := 0; i < 6; i++ {

		var kursi model.Kursi
		model.Db.Raw("SELECT * FROM kursi LEFT JOIN penayangan ON penayangan.audiotorium_id = kursi.audiotorium_id WHERE penayangan.id = ? LIMIT 1", i+1).Scan(&kursi)
		var seats []*model.Seat
		seats = append(seats, &model.Seat{KursiID: kursi.ID})
		tiket = append(tiket, model.Tiket{TotalHarga: 10000, UserID: 1, PenayanganID: uint(i + 1), Seat: seats})
	}
	model.Db.Create(tiket)
}
