package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gowaves/vendor_automation/internal/middlewares"
	userController "github.com/gowaves/vendor_automation/internal/modules/user/controllers"
)

func Routes(router *gin.Engine) {
	userController := userController.New()

	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsGuest())
	{
		//Register
		guestGroup.GET("/register", userController.Register)
		guestGroup.POST("/register", userController.HandleRegister)

		//Login
		guestGroup.GET("/login", userController.Login)
		guestGroup.POST("/login", userController.HandleLogin)
	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		//Logout
		router.POST("/logout", userController.HandleLogout)
	}

}
