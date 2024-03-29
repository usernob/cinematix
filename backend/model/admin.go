package model

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Nama      string         `json:"nama"`
	Email     string         `json:"email" gorm:"unique;not null;index"`
	Password  string         `json:"password"`
	Role      Role           `sql:"type:role" json:"role" gorm:"default:admin"`
	Avatar    *string        `json:"avatar"`
	Tiket     []*Tiket       `gorm:"foreignKey:SignedBy" json:"tiket,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
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

func GetAdminList() ([]Admin, error) {
	var admin []Admin
	res := Db.Where(&Admin{Role: AdminRole}).Find(&admin)
	if res.Error != nil {
		return nil, res.Error
	}
	return admin, nil
}

func CreateAdmin(nama string, email string, password string) (*Admin, error) {
	admin := Admin{
		Nama:     nama,
		Email:    email,
		Password: password,
		Role:     AdminRole,
	}
	res := Db.Create(&admin)
	if res.Error != nil {
		return nil, res.Error
	}
	return &admin, nil
}

func DeleteAdmin(id uint) error {
	res := Db.Where("id = ?", id).Unscoped().Delete(&Admin{})
	return res.Error
}
