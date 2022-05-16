package models

import (
	"database/sql"
	"time"
)

//OptionAccountLogs 账户日志表
type OptionAccountLogs struct {
	Id         int            `json:"id"`
	CurrencyId int            `json:"currency_id"`
	Type       int            `json:"type"`
	UserId     int            `json:"user_id"`
	FromId     int            `json:"from_id"` // 预留字段
	Before     float32        `json:"before"`  // 更改前
	Value      float32        `json:"value"`   // 更改值
	After      float32        `json:"after"`   // 更改后
	IsLock     int            `json:"is_lock"` // 是否锁定余额
	Memo       string         `json:"memo"`
	ExtraData  sql.NullString `json:"extra_data"` // 扩展数据
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}
