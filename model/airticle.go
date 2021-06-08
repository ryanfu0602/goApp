package model

type Article struct {
	ID       int    `json:"id"`
	CreateAT string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	Title    string `json:"title"`
	Author   string `json:"author"`
}

type ArticleUpdate struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
