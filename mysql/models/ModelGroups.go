package models

import "database/sql"

type ModelGroups struct {
	Id         string       `json:"id"`
	GroupName  string       `json:"group_name"`  // 组名称
	GroupField string       `json:"group_field"` // 组内字段
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}
