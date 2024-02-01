package model

import (
	"backend/pkg/passwordhash"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func Migrate() {
	db.AutoMigrate(&User{}, &Admin{}, &Film{}, &Genre{}, &Audiotorium{}, &Penayangan{}, &Tiket{}, &Kursi{}, &Seat{})

  hashedPassword, err := passwordhash.HashPassword(os.Getenv("SUPERADMIN_PASSWORD"))

  if err != nil {
    panic(err)
  }

  db.Create(&Admin{
    Nama: os.Getenv("SUPERADMIN_NAME"),
    Email: os.Getenv("SUPERADMIN_EMAIL"),
    Password: hashedPassword,
    Role: "superadmin",
  })
}

func SeedRefresh() {
	db.Migrator().DropTable(&User{}, &Admin{}, &Film{}, &Genre{}, &Penayangan{}, &Tiket{}, &Seat{}, &Kursi{}, &Audiotorium{})
	Migrate()
	Seeder()
}
