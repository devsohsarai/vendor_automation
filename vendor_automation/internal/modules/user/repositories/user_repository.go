package repositories

import (
	userModel "github.com/gowaves/vendor_automation/internal/modules/user/models"
	"github.com/gowaves/vendor_automation/pkg/database"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (userRepository *UserRepository) Create(user userModel.User) userModel.User {
	var newUser userModel.User
	userRepository.DB.Create(&user).Scan(&newUser)
	return newUser
}

func (userRepository *UserRepository) FindByEmail(email string) userModel.User {
	var user userModel.User
	userRepository.DB.First(&user, "email = ?", email)
	return user
}

func (userRepository *UserRepository) FindByID(id int) userModel.User {
	var user userModel.User
	userRepository.DB.First(&user, "id = ?", id)
	return user
}
