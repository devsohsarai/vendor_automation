package routes

import (
	"github.com/gin-gonic/gin"
	homeCtrl "github.com/gowaves/vendor_automation/internal/modules/home/controllers"
)

func Routes(router *gin.Engine) {
	homeController := homeCtrl.New()
	router.GET("/", homeController.Index)
}
