package models

import "database/sql"

type SubFund struct {
	Id                  string          `json:"id"`
	FundId              sql.NullInt64   `json:"fund_id"`
	UserId              sql.NullInt64   `json:"user_id"`
	SubTime             sql.NullTime    `json:"sub_time"` // 认购时间
	CreatedAt           sql.NullTime    `json:"created_at"`
	UpdatedAt           sql.NullTime    `json:"updated_at"`
	ShareAmount         sql.NullFloat64 `json:"share_amount"`           // 认购金额
	IsReturn            sql.NullInt64   `json:"is_return"`              // 是否退还本金
	Status              sql.NullInt64   `json:"status"`                 // 1.进行中 2.申请退款中,3.已退款 4已结束
	SurplusPeriod       sql.NullInt64   `json:"surplus_period"`         // 剩余未产生收益期数
	InterestGenNextTime sql.NullString  `json:"interest_gen_next_time"` // 利息生成时间
	FundAmount          sql.NullFloat64 `json:"fund_amount"`            // 产生的利息
}
