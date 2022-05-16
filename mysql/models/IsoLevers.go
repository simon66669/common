package models

import "time"

type IsoLevers struct {
	Id                 string    `json:"id"`
	Type               string    `json:"type"`                // 买卖类型:1.买入,2.卖出
	UserId             string    `json:"user_id"`             // 用户id
	CurrencyMatchId    string    `json:"currency_match_id"`   // 交易对id
	OriginPrice        float32   `json:"origin_price"`        // 原始价格
	FactPrice          float32   `json:"fact_price"`          // 开仓价格(点差处理之后)
	UpdatePrice        float32   `json:"update_price"`        // 当前价格
	TargetProfitPrice  float32   `json:"target_profit_price"` // 止盈价格
	StopLossPrice      float32   `json:"stop_loss_price"`     // 止亏价格
	Share              string    `json:"share"`               // 手数
	Number             float32   `json:"number"`              // 手数换算数量(非放大的)
	Multiple           string    `json:"multiple"`            // 倍数
	OriginCautionMoney float32   `json:"origin_caution_money"`
	CautionMoney       float32   `json:"caution_money"`   // 当前可用保证金
	FactProfits        float32   `json:"fact_profits"`    // 最终盈亏
	TradeFee           float32   `json:"trade_fee"`       // 交易手续费
	Overnight          float32   `json:"overnight"`       // 隔夜费率
	OvernightMoney     float32   `json:"overnight_money"` // 隔夜费金额
	Status             int       `json:"status"`          // 交易状态:0.挂单中,1.交易中,2.平仓中,3.已平仓,4.已撤单
	Settled            int       `json:"settled"`         // 结算状态:0.未结算,1.已结算
	CreatedAt          time.Time `json:"created_at"`      // 下单时间
	TransactedAt       time.Time `json:"transacted_at"`   // 交易时间
	UpdatedAt          time.Time `json:"updated_at"`      // 价格刷新时间(毫秒级)
	ClosedAt           time.Time `json:"closed_at"`       // 平仓时间
	ClosedType         int       `json:"closed_type"`     // 平仓类型:0.未知,1.市价,2.暴仓,3.止盈,4.止损,5.后台
	AgentPath          string    `json:"agent_path"`      // 用户代理商关系记录
	BoomPrice          float32   `json:"boom_price"`      // 强平价
}
