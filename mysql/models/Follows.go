package models

import "database/sql"

type Follows struct {
	Id        string        `json:"id"`
	UserId    sql.NullInt64 `json:"user_id"`
	Status    sql.NullInt64 `json:"status"` // 0停止跟随 1正在跟随
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}
