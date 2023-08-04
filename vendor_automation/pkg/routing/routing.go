package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/gowaves/vendor_automation/internal/providers/setroutes"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	setroutes.SetupRegisterRoutes(GetRouter())
}
