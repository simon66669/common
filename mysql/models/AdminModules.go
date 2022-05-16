package models

import "database/sql"

//AdminModules 系统模块表
type AdminModules struct {
	Id                string        `json:"id"`
	Name              string        `json:"name"`
	AdminModuleTypeId int           `json:"admin_module_type_id"`
	IsAddLog          sql.NullInt64 `json:"is_add_log"` // 是否添加日志，0不添加，1添加
}
