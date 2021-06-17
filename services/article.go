package services

import (
	"fmt"
	"goApp/dao"
	"goApp/model"

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

func UpdateArticleById(id int, article *model.Article) string {
	// res := "update id = " + strconv.Itoa(id)
	res := dao.UpdateArticleById(id, article)
	return res
}

func DeleteArticleById(id int) string {
	res := dao.DeleteArticleById(id)
	fmt.Println("res=", res)
	return res
}

func GetArticleList(title string, author string, country string, name string, age string, status string) *[]model.Article {
	res := dao.GetArticle(title, author, country, name, age, status)
	return res
}
