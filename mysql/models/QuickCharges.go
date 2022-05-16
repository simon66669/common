package models

import "database/sql"

type QuickCharges struct {
	Id               string         `json:"id"`
	CurrencyId       sql.NullInt64  `json:"currency_id"`
	Address          sql.NullString `json:"address"`
	CreatedAt        sql.NullTime   `json:"created_at"`
	UpdatedAt        sql.NullTime   `json:"updated_at"`
	CurrencyProtocId sql.NullInt64  `json:"currency_protoc_id"`
}
