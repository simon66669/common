package models

import "database/sql"

//MicroNumbers 秒合约数量
type MicroNumbers struct {
	Id         string       `json:"id"`
	CurrencyId string       `json:"currency_id"` // 币种id
	Number     float32      `json:"number"`      // 数量
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}
