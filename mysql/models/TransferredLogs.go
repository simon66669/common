package models

import "time"

//TransferredLogs 账户划转日志
type TransferredLogs struct {
	Id         string    `json:"id"`
	CurrencyId int       `json:"currency_id"`
	UserId     int       `json:"user_id"`
	From       string    `json:"from"`
	To         string    `json:"to"`
	Balance    float32   `json:"balance"` // 变动金额
	Memo       string    `json:"memo"`    // 备注
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
