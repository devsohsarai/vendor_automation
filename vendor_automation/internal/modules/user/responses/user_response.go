package responses

import (
	"fmt"

	UserModel "github.com/gowaves/vendor_automation/internal/modules/user/models"
)

type User struct {
	ID      uint
	Image   string
	Name    string
	Email   string
	Contact string
}

type Users struct {
	Data []User
}

func ToUser(user UserModel.User) User {
	return User{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Contact: user.Contact,
		Image:   fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
	}
}
