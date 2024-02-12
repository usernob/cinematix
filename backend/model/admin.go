package model

import (
	"database/sql/driver"
	"time"
)

type Admin struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `json:"nama"`
	Email     string    `json:"email" gorm:"unique;not null;index"`
	Password  string    `json:"password"`
	Role      Role      `sql:"type:role" json:"role" gorm:"default:admin"`
	Avatar    *string   `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role string

func (p *Role) Scan(value interface{}) error {
	*p = Role(value.(string))
	return nil
}

func (p Role) Value() (driver.Value, error) {
	return string(p), nil
}

const (
	SuperAdminRole Role = "superadmin"
	AdminRole      Role = "admin"
)

func GetAdminByEmail(email string) (*Admin, error) {
	var admin Admin
	res := Db.Where("email = ?", email).First(&admin)

	if res.Error != nil {
		return nil, res.Error
	}

	return &admin, nil
}

func GetAdminBy(admin Admin) (*Admin, error) {
	res := Db.Where(&admin).First(&admin)
	if res.Error != nil {
		return nil, res.Error
	}
	return &admin, nil
}

func CreateAdmin(nama string, email string, password string) (*Admin, error) {
	admin := Admin{
		Nama:     nama,
		Email:    email,
		Password: password,
	}
	res := Db.Create(&admin)
	if res.Error != nil {
		return nil, res.Error
	}
	return &admin, nil
}
