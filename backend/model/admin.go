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
	Role      Role           `sql:"type:role" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Role string

func (p *Role) Scan(value interface{}) error {
	*p = Role(value.([]byte))
	return nil
}

func (p Role) Value() (driver.Value, error) {
	return string(p), nil
}

const (
	SuperAdminRole Role = "superadmin"
	AdminRole      Role = "admin"
)
