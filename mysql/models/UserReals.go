package models

import (
	"database/sql"
	"time"
)

type UserReals struct {
	Id              string        `json:"id"`
	UserId          int           `json:"user_id"`
	Name            string        `json:"name"`
	CardId          string        `json:"card_id"`
	ReviewStatus    int           `json:"review_status"` // 认证状态
	FrontPic        string        `json:"front_pic"`
	ReversePic      string        `json:"reverse_pic"`
	HandPic         string        `json:"hand_pic"`
	CreatedAt       time.Time     `json:"created_at"` // 创建时间
	UpdatedAt       time.Time     `json:"updated_at"` // 确认时间
	IsCreateAddress sql.NullInt64 `json:"is_create_address"`
}
