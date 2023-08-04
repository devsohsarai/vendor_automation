package htmlparse

import "github.com/gin-gonic/gin"

func LoadHTML(router *gin.Engine) {
	//internal/modules/modulename/html/view.tmpl
	router.LoadHTMLGlob("internal/**/**/**/*tmpl")
}
