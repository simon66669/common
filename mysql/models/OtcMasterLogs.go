package models

import "database/sql"

//OtcMasterLogs 挂单操作日志
type OtcMasterLogs struct {
	Id           string       `json:"id"`
	MasterId     string       `json:"master_id"`     // 挂单id
	UserId       int          `json:"user_id"`       // 用户id
	OperatorType int          `json:"operator_type"` // 操作员类型:0.用户,1.系统,2.后台管理员
	Status       int          `json:"status"`        // 状态
	Memo         string       `json:"memo"`          // 说明
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}
