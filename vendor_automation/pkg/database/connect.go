package database

import (
	"fmt"
	"log"

	"github.com/gowaves/vendor_automation/pkg/cif"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	cfg := cif.Get()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cound not connect to the databae")
		return
	}
	DB = db
}
