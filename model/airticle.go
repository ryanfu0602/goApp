package model

import "gorm.io/datatypes"

type Article struct {
	ID      int            `json:"id"`
	Title   string         `json:"title"`
	Author  string         `json:"author"`
	Country string         `json:"country"`
	Other   datatypes.JSON `json:"other"`
}

type ArticleUpdate struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type ArticleAdd struct {
	Title   string         `json:"title"`
	Author  string         `json:"author"`
	Country string         `json:"country"`
	Other   datatypes.JSON `json:"other"`
}
