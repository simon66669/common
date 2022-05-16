package models

import "database/sql"

type SubProjects struct {
	Id              string          `json:"id"`
	ProjectId       sql.NullInt64   `json:"project_id"`
	BuyCurrencyId   sql.NullInt64   `json:"buy_currency_id"`  // 购买项目的币种
	BuyNumber       sql.NullFloat64 `json:"buy_number"`       // 购买金额
	PurchasedNumber sql.NullFloat64 `json:"purchased_number"` // 购买数量
	CurrencyMatchId sql.NullInt64   `json:"currency_match_id"`
	Price           sql.NullFloat64 `json:"price"`
	UserId          sql.NullInt64   `json:"user_id"`
	CreatedAt       sql.NullTime    `json:"created_at"`
	UpdatedAt       sql.NullTime    `json:"updated_at"`
}
