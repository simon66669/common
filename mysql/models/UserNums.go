package models

import "time"

//UserNums 用户相关资源统计，例如：直推、团队人数、业绩等
type UserNums struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`    // 用户id
	TypeName  string    `json:"type_name"`  // 类型名称
	Num       float32   `json:"num"`        // 数量
	Memo      string    `json:"memo"`       // 说明
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}
