package models

import "database/sql"

//OptionSeconds 微合约周期
type OptionSeconds struct {
	Id          string       `json:"id"`
	Seconds     string       `json:"seconds"`      // 秒数
	Status      int          `json:"status"`       // 状态:0.禁用,1.启用
	ProfitRatio float32      `json:"profit_ratio"` // 收益率
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}
