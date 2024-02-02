package model

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Nama      string         `json:"nama"`
	Email     string         `json:"email" gorm:"unique;not null;index"`
	Password  string         `json:"password"`
	Tiket     []*Tiket       `json:"tiket,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	res := Db.Where("email = ?", email).First(&user)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user not found")
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func CreateUser(nama string, email string, password string) (*User, error) {
	user := User{
		Nama:     nama,
		Email:    email,
		Password: password,
	}
	res := Db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
