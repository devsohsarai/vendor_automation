package seeder

import (
	"fmt"
	"log"

	articleModel "github.com/gowaves/vendor_automation/internal/modules/article/models"
	userModel "github.com/gowaves/vendor_automation/internal/modules/user/models"
	"github.com/gowaves/vendor_automation/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte("Secret"), 12)
	if err != nil {
		log.Fatal("hash password error")
		return
	}

	user := userModel.User{Name: "random name", Email: "randomuser@gmail.com", Password: string(hashPassword)}

	db.Create(&user) // pass pointer of data to Create

	log.Printf("User created successfully with email address %s\n", user.Email)

	for i := 1; i <= 10; i++ {
		article := articleModel.Article{Title: fmt.Sprintf("random title %d", i), Content: fmt.Sprintf("random content %d", i), UserID: user.ID}
		db.Create(&article)
		log.Printf("Article created successfully with title %s \n", article.Title)
	}

	log.Println("Seeder done ..")
}
