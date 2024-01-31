package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func Migrate() {
	db.AutoMigrate(&User{}, &Admin{}, &Film{}, &Genre{}, &Audiotorium{}, &Penayangan{}, &Tiket{}, &Kursi{}, &Seat{})
}

func SeedRefresh() {
	db.Migrator().DropTable(&User{}, &Admin{}, &Film{}, &Genre{}, &Penayangan{}, &Tiket{}, &Seat{}, &Kursi{}, &Audiotorium{})
	Migrate()
	Seeder()
}
