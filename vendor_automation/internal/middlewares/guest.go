package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	UserRepository "github.com/gowaves/vendor_automation/internal/modules/user/repositories"
	"github.com/gowaves/vendor_automation/pkg/sessions"
)

func IsGuest() gin.HandlerFunc {
	var userRepo = UserRepository.New()
	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)
		if user.ID != 0 {
			c.Redirect(http.StatusFound, "/")
			return
		}

		c.Next()

	}
}
