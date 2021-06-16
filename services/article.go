package services

import (
	"goApp/dao"
	"goApp/model"
	"strconv"

	"gorm.io/gorm"
)

func GetArticleById(id int) *model.Article {
	res := dao.GetArticleById(id)
	return res
}

func AddArticle(article *model.Article) *gorm.DB {
	res := dao.AddArticle(article)
	return res
}

func UpdateArticleById(id int, article *model.ArticleUpdate) string {
	// res := "update id = " + strconv.Itoa(id)
	res := dao.UpdateArticleById(id, article)
	return res
}

func DeleteArticleById(id int) string {
	res := "delete id = " + strconv.Itoa(id)
	return res
}
