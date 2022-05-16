package models

import "database/sql"

type Subscription struct {
	Id          int             `json:"id"`
	CurrencyId  sql.NullInt64   `json:"currency_id"`
	IssueNumber sql.NullInt64   `json:"issue_number"` // 最大发行数量
	Subscribed  sql.NullInt64   `json:"subscribed"`   // 已申购数量
	StartTime   sql.NullTime    `json:"start_time"`   // 申购开始时间
	FinishTime  sql.NullTime    `json:"finish_time"`  // 申购结束时间
	MarketTime  sql.NullTime    `json:"market_time"`  // 上市时间
	SubPrice    sql.NullFloat64 `json:"sub_price"`    // 申购价格
	MarketPrice sql.NullFloat64 `json:"market_price"` // 上市价格
	CreatedAt   sql.NullTime    `json:"created_at"`
	UpdatedAt   sql.NullTime    `json:"updated_at"`
	Status      sql.NullInt64   `json:"status"`  // 0未发币 1已发币
	IsShow      sql.NullInt64   `json:"is_show"` // 0不显示1显示
}
