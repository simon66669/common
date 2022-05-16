package models

import "database/sql"

type SubFundInterest struct {
	Id        int             `json:"id"`
	UserId    sql.NullInt64   `json:"user_id"`
	Interest  sql.NullFloat64 `json:"interest"` // 利息
	FundId    sql.NullInt64   `json:"fund_id"`
	SubId     sql.NullInt64   `json:"sub_id"`
	CreatedAt sql.NullTime    `json:"created_at"`
	UpdatedAt sql.NullTime    `json:"updated_at"`
}
