package models

import "database/sql"

type QuickChargeOrders struct {
	Id         string         `json:"id"`
	Uid        sql.NullInt64  `json:"uid"`
	Number     sql.NullString `json:"number"`
	ProofImg   sql.NullString `json:"proof_img"`
	Address    sql.NullString `json:"address"`
	UpdatedAt  sql.NullTime   `json:"updated_at"`
	CreatedAt  sql.NullTime   `json:"created_at"`
	ChargeId   sql.NullInt64  `json:"charge_id"`
	CurrencyId sql.NullInt64  `json:"currency_id"`
	Status     sql.NullInt64  `json:"status"`
}
