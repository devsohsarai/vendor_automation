package htmlparse

import (
	"github.com/gin-gonic/gin"
	"github.com/gowaves/vendor_automation/internal/providers/htmlglobal"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	data = htmlglobal.WithGlobalData(c, data)
	c.HTML(code, name, data)
}
