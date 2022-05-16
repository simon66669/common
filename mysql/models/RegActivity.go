package models

import "database/sql"

type RegActivity struct {
	Id         int             `json:"id"`
	IsOpen     sql.NullInt64   `json:"is_open"`
	CurrencyId sql.NullInt64   `json:"currency_id"`
	GiveNumber sql.NullFloat64 `json:"give_number"`
	CreatedAt  sql.NullTime    `json:"created_at"`
	UpdatedAt  sql.NullTime    `json:"updated_at"`
}
