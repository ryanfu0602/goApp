package services

import (
	"fmt"
	"goApp/dao"
	"goApp/model"
	"strconv"
)

func GetArticleById(id int) string {
	res := fmt.Sprintf("article id is %d", id)
	return res
}

func AddArticle(article *model.ArticleUpdate) *model.ArticleUpdate {
	res := dao.AddArticle(article)
	return res
}

func UpdateArticleById(id int, article *model.ArticleUpdate) string {
	res := "update id = " + strconv.Itoa(id)
	return res
}

func DeleteArticleById(id int) string {
	res := "delete id = " + strconv.Itoa(id)
	return res
}
