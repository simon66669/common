package models

import "database/sql"

type RegActivityLists struct {
	Id         int             `json:"id"`
	UserId     sql.NullInt64   `json:"user_id"`
	Amount     sql.NullFloat64 `json:"amount"`
	CurrencyId sql.NullInt64   `json:"currency_id"`
	CreatedAt  sql.NullTime    `json:"created_at"`
	UpdatedAt  sql.NullTime    `json:"updated_at"`
}
