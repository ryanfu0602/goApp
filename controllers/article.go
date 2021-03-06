package controllers

import (
	"fmt"
	"goApp/model"
	"goApp/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetArticleById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}
	res := services.GetArticleById(id)
	return c.JSON(http.StatusOK, res)
}

func GetArticleList(c echo.Context) error {
	title := c.QueryParam("title")
	author := c.QueryParam("author")
	country := c.QueryParam("country")
	name := c.QueryParam("name")
	age := c.QueryParam("age")
	status := c.QueryParam("status")
	println(title, author, country, name, age)

	res := services.GetArticleList(title, author, country, name, age, status)
	return c.JSON(http.StatusOK, res)
}

func AddArticle(c echo.Context) error {
	article := new(model.Article)
	if err := c.Bind(article); err != nil {
		return err
	}
	res := services.AddArticle(article)
	fmt.Println("error=", res.Error)
	return c.JSON(http.StatusCreated, "OK")
}

func UpdateArticleById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}
	article := new(model.Article)
	if err := c.Bind(article); err != nil {
		return err
	}
	res := services.UpdateArticleById(id, article)
	return c.JSON(http.StatusOK, res)
}

func DeleteArticleById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}
	res := services.DeleteArticleById(id)
	return c.JSON(http.StatusOK, res)
}
