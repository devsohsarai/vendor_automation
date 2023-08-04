package migration

import (
	"fmt"
	"log"

	articleModel "github.com/gowaves/vendor_automation/internal/modules/article/models"
	usermodel "github.com/gowaves/vendor_automation/internal/modules/user/models"

	"github.com/gowaves/vendor_automation/pkg/database"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&articleModel.Article{}, &usermodel.User{})
	if err != nil {
		log.Fatal("Can not migrate, please coordinate with admin!")
		return
	}

	fmt.Println("Migration done.....")

}
