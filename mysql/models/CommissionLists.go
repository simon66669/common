package models

import "database/sql"

type CommissionLists struct {
	Id          string          `json:"id"`
	FromUserId  sql.NullString  `json:"from_user_id"`
	ToUserId    sql.NullInt64   `json:"to_user_id"`
	OrderAmount sql.NullFloat64 `json:"order_amount"`
	CurrencyId  sql.NullInt64   `json:"currency_id"`
	CreatedAt   sql.NullTime    `json:"created_at"`
	UpdatedAt   sql.NullTime    `json:"updated_at"`
	Status      sql.NullInt64   `json:"status"` // 1.未处理 2.不返佣 3.已返佣
	ExcTime     sql.NullTime    `json:"exc_time"`
	ExcAmount   sql.NullFloat64 `json:"exc_amount"`
}
