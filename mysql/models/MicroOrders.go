package models

import "database/sql"

//MicroOrders 微合约订单表
type MicroOrders struct {
	Id              string         `json:"id"`
	UserId          string         `json:"user_id"`           // 用户id
	MatchId         string         `json:"match_id"`          // 交易对id
	CurrencyId      string         `json:"currency_id"`       // 支付的币种
	Type            int            `json:"type"`              // 买卖类型1.买涨,2.买跌
	IsInsurance     int            `json:"is_insurance"`      // 订单险种:0.无,1.正向，2反向。
	Seconds         string         `json:"seconds"`           // 秒数
	Number          float32        `json:"number"`            // 下单数量
	OpenPrice       float32        `json:"open_price"`        // 开仓价
	EndPrice        float32        `json:"end_price"`         // 收盘价
	Fee             float32        `json:"fee"`               // 手续费
	ProfitRatio     float32        `json:"profit_ratio"`      // 收益率
	FactProfits     float32        `json:"fact_profits"`      // 最终收益
	Status          int            `json:"status"`            // 状态:1.交易中,2.平仓中,3.已平仓
	PreProfitResult int            `json:"pre_profit_result"` // 预设盈利状态:-1.亏损,0.未设置,1.盈利
	ProfitResult    int            `json:"profit_result"`     // 盈利结果:-1.亏损,0.平,1.盈利
	CreatedAt       sql.NullTime   `json:"created_at"`        // 提交日期
	UpdatedAt       sql.NullTime   `json:"updated_at"`        // 更新日期
	HandledAt       sql.NullTime   `json:"handled_at"`        // 平仓时间
	CompleteAt      sql.NullTime   `json:"complete_at"`       // 完成时间
	AgentPath       sql.NullString `json:"agent_path"`        // 代理商关系
	Settled         int            `json:"settled"`           // 是否已给代理商结算1已结算
}
