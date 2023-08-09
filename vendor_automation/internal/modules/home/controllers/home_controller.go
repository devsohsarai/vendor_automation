package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ArticleService "github.com/gowaves/vendor_automation/internal/modules/article/services"
	"github.com/gowaves/vendor_automation/pkg/htmlparse"
)

type Controller struct {
	//articleService ArticleService.ArticleServiceInterface
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	htmlparse.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":      "Welcome to the vendor automation",
		"featured":   controller.articleService.GetFeaturedArticles(),
		"stories":    controller.articleService.GetStoriesArticles(),
		"ActivePage": "home",
	})
}

func (controller *Controller) About(c *gin.Context) {
	htmlparse.Render(c, http.StatusOK, "modules/home/html/about", gin.H{
		"title":      "About us",
		"ActivePage": "about",
	})
}

func (controller *Controller) Services(c *gin.Context) {
	htmlparse.Render(c, http.StatusOK, "modules/home/html/services", gin.H{
		"title":      "Services",
		"ActivePage": "services",
	})
}

func (controller *Controller) Contact(c *gin.Context) {
	htmlparse.Render(c, http.StatusOK, "modules/home/html/contact", gin.H{
		"title":      "Conatact",
		"ActivePage": "contact",
	})
}
