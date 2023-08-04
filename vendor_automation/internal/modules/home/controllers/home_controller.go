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
		"title":    "Welcome to the eSeller360",
		"featured": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),
	})
}
