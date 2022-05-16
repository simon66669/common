package models

import "database/sql"

//AdminModuleTypes 管理员权限分类表
type AdminModuleTypes struct {
	Id   int            `json:"id"`
	Name sql.NullString `json:"name"`
}
