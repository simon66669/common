package models

import "database/sql"

type OptionPeriods struct {
	Id                string          `json:"id"`
	Period            sql.NullInt64   `json:"period"`
	SecondsId         sql.NullInt64   `json:"seconds_id"`          // seconds表id
	Status            sql.NullInt64   `json:"status"`              // 0未开始 1正在运行 2结算中 3已结束
	Result            sql.NullInt64   `json:"result"`              // 0未开奖1涨2跌
	ResultTime        sql.NullTime    `json:"result_time"`         // 结束时间
	StartTime         sql.NullTime    `json:"start_time"`          // 开始时间
	CreatedTime       sql.NullTime    `json:"created_time"`        // 创建时间
	CapStartTimestamp sql.NullInt64   `json:"cap_start_timestamp"` // 开始时间戳
	CapEndTimestamp   sql.NullInt64   `json:"cap_end_timestamp"`   // 结束时间戳
	CurrencyMatchId   sql.NullInt64   `json:"currency_match_id"`
	OpenPrice         sql.NullFloat64 `json:"open_price"`  // 开仓价格
	ClosePrice        sql.NullFloat64 `json:"close_price"` // 平仓价格
}
