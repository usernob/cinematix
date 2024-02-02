package seeder

import (
	"backend/model"
	"backend/pkg/passwordhash"
)

func Seed() {
	pass, err := passwordhash.HashPassword("password")
	if err != nil {
		panic(err)
	}

	model.Db.Create(&model.User{
		Nama:     "user1",
		Email:    "user1@gmail.com",
		Password: pass,
	})

  seedGenre()
  seedFilmAndPenyangan()
}

