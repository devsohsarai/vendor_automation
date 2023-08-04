package repositories

import ArticleModel "github.com/gowaves/vendor_automation/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
	Find(id int) ArticleModel.Article
	Create(article ArticleModel.Article) ArticleModel.Article
}
