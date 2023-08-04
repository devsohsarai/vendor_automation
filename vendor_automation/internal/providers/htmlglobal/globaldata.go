package htmlglobal

import (
	"github.com/gin-gonic/gin"
	"github.com/gowaves/vendor_automation/internal/modules/user/helpers"
	"github.com/gowaves/vendor_automation/pkg/converters"
	"github.com/gowaves/vendor_automation/pkg/sessions"
	"github.com/spf13/viper"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("App.Name")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(c, "errors"))
	data["OLD"] = converters.StringToUrlValues(sessions.Flash(c, "old"))
	data["AUTH"] = helpers.Auth(c)
	return data
}
