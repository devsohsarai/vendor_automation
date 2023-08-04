package repositories

import (
	ArticleModel "github.com/gowaves/vendor_automation/internal/modules/article/models"
	"github.com/gowaves/vendor_automation/pkg/database"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article
	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)
	return articles
}

func (articleRepository *ArticleRepository) Find(id int) ArticleModel.Article {
	var article ArticleModel.Article
	articleRepository.DB.Joins("User").Find(&article, id)
	return article
}

func (articleRepository *ArticleRepository) Create(article ArticleModel.Article) ArticleModel.Article {
	var newArticle ArticleModel.Article
	articleRepository.DB.Create(&article).Scan(&newArticle)
	return newArticle

}
