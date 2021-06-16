package dao

import (
	"fmt"
	"goApp/model"
	"goApp/utils"

	"gorm.io/gorm"
)

func AddArticle(article *model.Article) *gorm.DB {
	db := utils.GetDB()
	result := db.Create(article)

	return result
}

func GetArticleById(id int) *model.Article {
	db := utils.GetDB()
	var result model.Article
	db.Raw("SELECT * FROM article WHERE id = ?", id).Scan(&result)
	fmt.Println(result)
	return &result
}

func UpdateArticleById(id int, article *model.ArticleUpdate) string {
	db := utils.GetDB()
	fmt.Println(article.Title)
	fmt.Println(article.Author)
	reslut := db.Exec("UPDATE article SET title=?, author=? WHERE id = ?", article.Title, article.Author, id)
	fmt.Println(reslut.RowsAffected)
	return "ok"
}
