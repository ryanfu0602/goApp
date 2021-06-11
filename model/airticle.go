package model

type Article struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type ArticleUpdate struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type ArticleAdd struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
