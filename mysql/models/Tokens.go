package models

import (
	"database/sql"
	"time"
)

//Tokens token表
type Tokens struct {
	Id        int           `json:"id"`
	UserId    int           `json:"user_id"`
	Token     string        `json:"token"`
	TimeoutAt time.Time     `json:"timeout_at"` // 过期时间
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Follower  sql.NullInt64 `json:"follower"`
}
