package services

import (
	"errors"

	userModel "github.com/gowaves/vendor_automation/internal/modules/user/models"
	UserRepository "github.com/gowaves/vendor_automation/internal/modules/user/repositories"
	"github.com/gowaves/vendor_automation/internal/modules/user/requests/auth"
	UserResponse "github.com/gowaves/vendor_automation/internal/modules/user/responses"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user userModel.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error hashing the password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)

	newuser := userService.userRepository.Create(user)
	if newuser.ID == 0 {
		return response, errors.New("erorr on creating the user")
	}

	return UserResponse.ToUser(newuser), nil
}

func (userService *UserService) CheckUserExists(email string) bool {
	user := userService.userRepository.FindByEmail(email)
	return user.ID != 0
}

func (userService *UserService) HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error) {
	var response UserResponse.User
	exitsUser := userService.userRepository.FindByEmail(request.Email)

	if exitsUser.ID == 0 {
		return response, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(exitsUser.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New("invalid credentials")
	}

	return UserResponse.ToUser(exitsUser), nil
}
