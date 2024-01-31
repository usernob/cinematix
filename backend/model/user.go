package model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Email    string
	Password string
	Tiket    []*Tiket
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	res := db.Where("email = ?", email).First(&user)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user not found")
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func CreateUser(nama string, email string, password string) error {
	res := db.Create(&User{
    Nama:     nama,
    Email:    email,
    Password: password,  	
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
