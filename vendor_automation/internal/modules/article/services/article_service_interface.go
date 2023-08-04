package services

import (
	"github.com/gowaves/vendor_automation/internal/modules/article/requests/articles"
	ArticleResponse "github.com/gowaves/vendor_automation/internal/modules/article/responses"
	UserResponse "github.com/gowaves/vendor_automation/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error)
}
