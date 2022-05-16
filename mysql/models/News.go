package models

import (
	"database/sql"
	"time"
)

//News 新闻表
type News struct {
	Id               string         `json:"id"`
	CategoryId       int            `json:"category_id"`
	LangId           int            `json:"lang_id"` // 语言代码
	Title            string         `json:"title"`   // 标题
	Display          int            `json:"display"`
	Author           string         `json:"author"` // 发布人
	Keywords         string         `json:"keywords"`
	Desc             string         `json:"desc"`
	Content          string         `json:"content"` // 内容
	Views            int            `json:"views"`   // 浏览量
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	Logo             string         `json:"logo"`              // 缩略图
	Banner           string         `json:"banner"`            // banner
	JumpSubscription sql.NullInt64  `json:"jump_subscription"` // 是否跳转订阅页
	WhitePaper       sql.NullString `json:"white_paper"`       // 白皮书地址
	Link             sql.NullString `json:"link"`              // 跳转链接
}
