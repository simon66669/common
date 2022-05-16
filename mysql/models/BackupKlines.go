package models

import "database/sql"

type BackupKlines struct {
	Id              string         `json:"id"`
	Period          sql.NullString `json:"period"`
	CurrencyMatchId sql.NullInt64  `json:"currency_match_id"`
	KlineData       sql.NullString `json:"kline_data"`
	UpdatedAt       sql.NullTime   `json:"updated_at"`
	CreatedAt       sql.NullTime   `json:"created_at"`
}
