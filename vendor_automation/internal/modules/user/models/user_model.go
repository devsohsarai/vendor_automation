package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(191)"`
	Email     string `gorm:"type:varchar(320)"`
	Password  string `gorm:"type:varchar(255)"`
	Contact   string `gorm:"type:varchar(20)"`
	IsAdmin   bool   // Default value will be false
	CompanyID uint16
	Company   Company `gorm:"foreignKey:CompanyID"`
}

type Company struct {
	gorm.Model
	ID      uint16 `gorm:"primaryKey"`
	Name    string `gorm:"type:varchar(191)"`
	Email   string `gorm:"type:varchar(320)"`
	Contact string `gorm:"type:varchar(20)"`
	Address string `gorm:"type:text"`
	Status  bool
	Users   []User // HasMany association
}
