package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gowaves/vendor_automation/internal/modules/article/requests/articles"
	ArticleService "github.com/gowaves/vendor_automation/internal/modules/article/services"
	"github.com/gowaves/vendor_automation/internal/modules/user/helpers"
	"github.com/gowaves/vendor_automation/pkg/converters"
	"github.com/gowaves/vendor_automation/pkg/errors"
	"github.com/gowaves/vendor_automation/pkg/htmlparse"
	"github.com/gowaves/vendor_automation/pkg/old"
	"github.com/gowaves/vendor_automation/pkg/sessions"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	//Get the article ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		htmlparse.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	//Find the article from the database
	article, err := controller.articleService.Find(id)

	// If the article is not found show the error page
	if err != nil {
		htmlparse.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	// If article found render article template
	htmlparse.Render(c, http.StatusOK, "modules/article/html/show", gin.H{"title": "Show Article", "article": article})
}

func (controller *Controller) Create(c *gin.Context) {
	htmlparse.Render(c, http.StatusOK, "modules/article/html/create", gin.H{"title": "Create Article"})
}

func (controller *Controller) Store(c *gin.Context) {
	// validate the request
	var request articles.StoreRequest

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}
	user := helpers.Auth(c)

	//Create the article
	article, err := controller.articleService.StoreAsUser(request, user)

	//Check if there is any error on the article creartion
	if err != nil {
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))

}
