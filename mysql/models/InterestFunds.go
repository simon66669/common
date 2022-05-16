package models

import "database/sql"

type InterestFunds struct {
	Id        string          `json:"id"`
	FundId    sql.NullInt64   `json:"fund_id"`
	UserId    sql.NullInt64   `json:"user_id"`
	Interest  sql.NullFloat64 `json:"interest"`
	CreatedAt sql.NullTime    `json:"created_at"`
	SubId     sql.NullInt64   `json:"sub_id"`
	PeriodsId sql.NullInt64   `json:"periods_id"`
}
