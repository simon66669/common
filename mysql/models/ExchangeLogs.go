package models

import "database/sql"

type ExchangeLogs struct {
	Id              string          `json:"id"`
	UserId          sql.NullInt64   `json:"user_id"`
	BaseCurrencyId  sql.NullInt64   `json:"base_currency_id"`
	QuoteCurrencyId sql.NullInt64   `json:"quote_currency_id"`
	LastPrice       sql.NullFloat64 `json:"last_price"`
	Type            sql.NullInt64   `json:"type"`
	Amount          sql.NullFloat64 `json:"amount"` // 兑换数量
	Number          sql.NullFloat64 `json:"number"` // 已兑换数量
	CreatedAt       sql.NullTime    `json:"created_at"`
	UpdatedAt       sql.NullTime    `json:"updated_at"`
}
