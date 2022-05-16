package models

import "database/sql"

type ProjectRobots struct {
	Id              string          `json:"id"`
	CurrencyMatchId sql.NullInt64   `json:"currency_match_id"`
	DefaultChange   sql.NullFloat64 `json:"default_change"`
	OpenPrice       sql.NullFloat64 `json:"open_price"`
	CreatedAt       sql.NullTime    `json:"created_at"`
	UpdatedAt       sql.NullTime    `json:"updated_at"`
	ProcessId       sql.NullInt64   `json:"process_id"`
	CurrentPrice    sql.NullFloat64 `json:"current_price"` // 当前价
	PushPrice       sql.NullFloat64 `json:"push_price"`
}
