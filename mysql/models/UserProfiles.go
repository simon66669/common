package models

import "database/sql"

//UserProfiles 用户资料表
type UserProfiles struct {
	Id        int          `json:"id"`
	UserId    int          `json:"user_id"`
	Nickname  string       `json:"nickname"` // 昵称
	Realname  string       `json:"realname"` // 真实姓名
	Avatar    string       `json:"avatar"`   // 头像
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
