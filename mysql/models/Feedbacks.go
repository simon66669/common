package models

import (
	"database/sql"
	"time"
)

//Feedbacks 用户反馈表
type Feedbacks struct {
	Id        string         `json:"id"`
	UserId    int            `json:"user_id"`    // 用户id
	TypeId    int            `json:"type_id"`    // 反馈类型id
	Title     string         `json:"title"`      // 标题
	Content   string         `json:"content"`    // 反馈内容
	Reply     sql.NullString `json:"reply"`      // 后台回复
	IsDisplay int            `json:"is_display"` // 是否展示,默认0不展示,1为展示
	IsReplied int            `json:"is_replied"` // 是否已回复,默认0未回复,1已回复
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
