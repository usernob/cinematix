package seeder

import (
	"backend/model"
	"fmt"
)

func seedCinema() {
	var kursis []model.Kursi

	var abjd string = "ABCDEFGHI"

	for j := 1; j < 3; j++ {
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
