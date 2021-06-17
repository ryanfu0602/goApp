package model

import "gorm.io/datatypes"

type Article struct {
	Title   string         `json:"title"`
	Author  string         `json:"author"`
	Country string         `json:"country"`
	Other   datatypes.JSON `json:"other"`
	Status  string         `json:"status"`
}
