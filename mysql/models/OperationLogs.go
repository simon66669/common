package models

import "time"

//OperationLogs 后台访问日志
type OperationLogs struct {
	Id          string    `json:"id"`
	AdminId     int       `json:"admin_id"`     // 后台登录用户id
	Method      string    `json:"method"`       // 方法
	Ip          string    `json:"ip"`           // 地址
	RequestPath string    `json:"request_path"` // 请求路径
	Data        string    `json:"data"`         // 参数
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
}
