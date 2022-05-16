package models

import "database/sql"

type FundsPeriods struct {
	Id            string          `json:"id"`
	FundId        sql.NullInt64   `json:"fund_id"`
	PeriodsNumber sql.NullString  `json:"periods_number"`  // 期数编号
	Status        sql.NullInt64   `json:"status"`          // 1.未发放 2.已发放
	GrantInterest sql.NullFloat64 `json:"grant_interest"`  // 当期总发放利息
	GrantTime     sql.NullTime    `json:"grant_time"`      // 发放时间
	GrantSendTime sql.NullTime    `json:"grant_send_time"` // 该期发放日
	EachDividend  sql.NullFloat64 `json:"each_dividend"`   // 当期派息百分比
	UpdatedAt     sql.NullTime    `json:"updated_at"`
	IsFy          sql.NullInt64   `json:"is_fy"` // 是否反佣
}
