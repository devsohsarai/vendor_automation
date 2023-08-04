package setroutes

import (
	"github.com/gin-gonic/gin"
	articleRoutes "github.com/gowaves/vendor_automation/internal/modules/article/routes"
	homeRoutes "github.com/gowaves/vendor_automation/internal/modules/home/routes"
	userRoutes "github.com/gowaves/vendor_automation/internal/modules/user/routes"
)

func SetupRegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)
}
