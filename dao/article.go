package dao

import (
	"fmt"
	"goApp/model"
	"goApp/utils"

	"gorm.io/gorm"
)

func AddArticle(article *model.ArticleAdd) *gorm.DB {
	db := utils.GetDB()

	newArticle := new(model.Article)
	newArticle.Title = article.Title
	newArticle.Author = article.Author

	result := db.Create(&newArticle)

	return result
}

func GetArticleById(id int) *model.Article {
	db := utils.GetDB()
	var result model.Article
	db.Raw("SELECT * FROM article WHERE id = ?", id).Scan(&result)
	fmt.Println(result)
	return &result
}
