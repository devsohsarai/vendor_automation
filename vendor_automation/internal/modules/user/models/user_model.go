package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"varchar:191"`
	Email     string `gorm:"varchar:191"`
	Password  string `gorm:"varchar:191"`
	Contact   string `gorm:"varchar:20"`
	IsAdmin   bool   // Default value will be false
	CompanyID uint16
	Company   Company `gorm:"foreignKey:CompanyID"`
}

type Company struct {
	gorm.Model
	ID      uint16 `gorm:"primaryKey"`
	Name    string `gorm:"varchar:250"`
	Email   string `gorm:"varchar:250"`
	Contact string `gorm:"varchar:20"`
	Address string `gorm:"type:text"`
	Status  bool
	Users   []User // HasMany association
}
