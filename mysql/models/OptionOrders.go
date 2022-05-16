package models

import "database/sql"

type OptionOrders struct {
	Id           string          `json:"id"`
	UserId       sql.NullInt64   `json:"user_id"`
	PeriodId     sql.NullInt64   `json:"period_id"`
	CreatedAt    sql.NullTime    `json:"created_at"`
	UpdatedAt    sql.NullTime    `json:"updated_at"`
	CurrencyId   sql.NullInt64   `json:"currency_id"` // 下单币种id
	MatchId      sql.NullInt64   `json:"match_id"`    // 交易对
	Number       sql.NullFloat64 `json:"number"`      // 数量
	Fee          sql.NullFloat64 `json:"fee"`         // 手续费
	Status       sql.NullInt64   `json:"status"`      // 1交易中2平仓中3已平仓
	NumberId     sql.NullInt64   `json:"number_id"`   // 数量id
	HandleTime   sql.NullTime    `json:"handle_time"`
	Result       sql.NullInt64   `json:"result"`        // 结果
	Type         sql.NullInt64   `json:"type"`          // 1看涨 2看跌
	ResultAmount sql.NullFloat64 `json:"result_amount"` // 最终盈亏
	SecondId     sql.NullInt64   `json:"second_id"`
}
