package models

import "time"

//NewsCategories 新闻分类表
type NewsCategories struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Sorts     int       `json:"sorts"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
