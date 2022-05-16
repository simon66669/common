package models

import "database/sql"

type LeverTransaction struct {
	Id                 string        `json:"id"`
	Type               string        `json:"type"`                 // 买卖类型:1.买入,2.卖出
	UserId             string        `json:"user_id"`              // 用户id
	CurrencyMatchId    string        `json:"currency_match_id"`    // 交易对id
	BaseCurrencyId     string        `json:"base_currency_id"`     // 交易币id
	QuoteCurrencyId    string        `json:"quote_currency_id"`    // 计价币id
	OriginPrice        float32       `json:"origin_price"`         // 原始价格
	Price              float32       `json:"price"`                // 开仓价格(点差处理之后)
	UpdatePrice        float32       `json:"update_price"`         // 当前价格
	TargetProfitPrice  float32       `json:"target_profit_price"`  // 止盈价格
	StopLossPrice      float32       `json:"stop_loss_price"`      // 止亏价格
	Share              string        `json:"share"`                // 手数
	Number             float32       `json:"number"`               // 手数换算数量(非放大的)
	Multiple           int           `json:"multiple"`             // 倍数
	OriginCautionMoney float32       `json:"origin_caution_money"` // 初始保证金
	CautionMoney       float32       `json:"caution_money"`        // 当前可用保证金
	FactProfits        float32       `json:"fact_profits"`         // 最终盈亏
	TradeFee           float32       `json:"trade_fee"`            // 交易手续费
	Overnight          float32       `json:"overnight"`            // 隔夜费率,百分比
	OvernightMoney     float32       `json:"overnight_money"`      // 隔夜费金额
	Status             int           `json:"status"`               // 交易状态:0.挂单中,1.交易中,2.平仓中,3.已平仓,4.已撤单
	Settled            int           `json:"settled"`              // 结算状态:0.未结算,1.已结算
	CreateTime         int           `json:"create_time"`          // 下单时间
	TransactionTime    float32       `json:"transaction_time"`     // 交易时间
	UpdateTime         float32       `json:"update_time"`          // 价格刷新时间(毫秒级)
	HandleTime         float32       `json:"handle_time"`          // 平仓时间
	CompleteTime       float32       `json:"complete_time"`        // 完成时间
	ClosedType         sql.NullInt64 `json:"closed_type"`          // 平仓类型:0.未知,1.市价,2.暴仓,3.止盈,4.止损,5.后台
	AgentPath          string        `json:"agent_path"`           // 用户代理商关系记录
}
