package models

import "database/sql"

type ApplyRefund struct {
	Id        int           `json:"id"`
	SubId     sql.NullInt64 `json:"sub_id"`
	FundId    sql.NullInt64 `json:"fund_id"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	UserId    sql.NullInt64 `json:"user_id"`
	Status    sql.NullInt64 `json:"status"` // 1.申请中 2.已驳回 3.已退款
}
