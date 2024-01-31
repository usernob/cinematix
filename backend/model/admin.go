package model

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type Admin struct {
  gorm.Model
  Nama string
  Email string
  Password string
  Role Role `sql:"type:role"`  
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
  AdminRole Role = "admin"
)

