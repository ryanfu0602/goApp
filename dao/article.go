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

func UpdateArticleById(id int, article *model.Article) string {
	db := utils.GetDB()
	fmt.Println(article.Title)
	fmt.Println(article.Author)
	reslut := db.Exec("UPDATE article SET title=?, author=? WHERE id = ?", article.Title, article.Author, id)
	fmt.Println(reslut.RowsAffected)
	return fmt.Sprintf("%d Rows Affected", reslut.RowsAffected)
}

func GetArticle(title string, author string, country string, name string, age string, status string) *[]model.Article {
	db := utils.GetDB()
	var result []model.Article
	sql := "SELECT * FROM  article WHERE 1=1 "
	if title != "" {
		sql += fmt.Sprintf("AND title = '%s'", title)
	}
	if author != "" {
		sql += fmt.Sprintf("AND author = '%s'", author)
	}
	if country != "" {
		sql += fmt.Sprintf("AND country = '%s'", country)
	}
	if status != "" {
		sql += fmt.Sprintf("AND status = '%s'", status)
	}
	if name != "" {
		sql += fmt.Sprintf("AND other->'$.name' = '%s'", name)
	}
	if age != "" {
		sql += fmt.Sprintf("AND other->'$.age' = %s", age)
	}
	db.Raw(sql).Scan(&result)
	fmt.Println(result)
	return &result
}

func DeleteArticleById(id int) string {
	db := utils.GetDB()
	reslut := db.Exec("UPDATE article SET status = 0,delete_at = current_timestamp WHERE id = ?", id)
	fmt.Println(reslut.RowsAffected)
	fmt.Println(reslut.Error)
	return fmt.Sprintf("%d Rows Affected", reslut.RowsAffected)
}
