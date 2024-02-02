package model

import (
	"backend/pkg/passwordhash"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func Migrate() {
  Db.SetupJoinTable(&Film{}, "Genres", &FilmGenre{})
	Db.AutoMigrate(&User{}, &Admin{}, &Film{}, &Genre{}, &Audiotorium{}, &Penayangan{}, &Tiket{}, &Kursi{}, &Seat{})

	hashedPassword, err := passwordhash.HashPassword(os.Getenv("SUPERADMIN_PASSWORD"))

	if err != nil {
		panic(err)
	}

	Db.Create(&Admin{
		Nama:     os.Getenv("SUPERADMIN_NAME"),
		Email:    os.Getenv("SUPERADMIN_EMAIL"),
		Password: hashedPassword,
		Role:     "superadmin",
	})
}

func Refresh() {
	Db.Migrator().DropTable(&User{}, &Admin{}, &Film{}, &Genre{}, &FilmGenre{}, &Penayangan{}, &Tiket{}, &Seat{}, &Kursi{}, &Audiotorium{})
}
