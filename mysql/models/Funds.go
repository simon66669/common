package models

import "database/sql"

type Funds struct {
	Id                 string          `json:"id"`
	Title              sql.NullString  `json:"title"`              // 产品标题
	CurrencyId         sql.NullInt64   `json:"currency_id"`        // 币种id
	LockDividendDays   sql.NullInt64   `json:"lock_dividend_days"` // 锁仓天数
	CreatedAt          sql.NullTime    `json:"created_at"`
	UpdatedAt          sql.NullTime    `json:"updated_at"`
	IsShow             sql.NullInt64   `json:"is_show"` // 是否显示
	Status             sql.NullInt64   `json:"status"`
	StaringSubAmount   sql.NullFloat64 `json:"staring_sub_amount"`  // 起购金额
	MaxSubAmount       sql.NullFloat64 `json:"max_sub_amount"`      // 最大下单金额
	LiquidatedDamages  sql.NullFloat64 `json:"liquidated_damages"`  // 提前赎回违约金
	DividendPercentage sql.NullFloat64 `json:"dividend_percentage"` // 每期派息百分比
	Image              sql.NullString  `json:"image"`
	Desc               sql.NullString  `json:"desc"` // 简介
}
