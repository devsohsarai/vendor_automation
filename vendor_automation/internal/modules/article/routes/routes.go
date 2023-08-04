package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gowaves/vendor_automation/internal/middlewares"
	articleCtrl "github.com/gowaves/vendor_automation/internal/modules/article/controllers"
)

func Routes(router *gin.Engine) {
	articleController := articleCtrl.New()

	router.GET("/articles/:id", articleController.Show)
	authGroup := router.Group("/articles")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/create", articleController.Create)
		authGroup.POST("/store", articleController.Store)
	}

}
