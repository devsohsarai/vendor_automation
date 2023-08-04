package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	UserRepository "github.com/gowaves/vendor_automation/internal/modules/user/repositories"
	UserResponse "github.com/gowaves/vendor_automation/internal/modules/user/responses"
	"github.com/gowaves/vendor_automation/pkg/sessions"
)

func Auth(c *gin.Context) UserResponse.User {
	var response UserResponse.User
	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)

	var userRepo = UserRepository.New()
	user := userRepo.FindByID(userID)

	if user.ID == 0 {
		return response
	}
	return UserResponse.ToUser(user)
}
