package models

import (
	"database/sql"
	"time"
)

//TxNew 交易买入表
type TxNew struct {
	Id              int             `json:"id"`
	Type            sql.NullInt64   `json:"type"`              // 1买  2卖
	CurrencyMatchId int             `json:"currency_match_id"` // 交易对id
	Price           float32         `json:"price"`             // 相对价格
	Number          float32         `json:"number"`
	Amount          float32         `json:"amount"`  // 已成交额
	Fee             float32         `json:"fee"`     // 手续费
	UserId          int             `json:"user_id"` // 用户id
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	QuoteCurrencyId sql.NullInt64   `json:"quote_currency_id"`
	BaseCurrencyId  sql.NullInt64   `json:"base_currency_id"`
	QuotePrice      sql.NullFloat64 `json:"quote_price"`
	BasePrice       sql.NullFloat64 `json:"base_price"`
	InviteCode      sql.NullString  `json:"invite_code"` // 邀请码
}
