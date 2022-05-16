package models

import "time"

//TxHistories 历史交易记录表
type TxHistories struct {
	Id               int       `json:"id"`
	CurrencyMatchId  int       `json:"currency_match_id"` // 交易对id
	Price            float32   `json:"price"`             // 价格
	Number           float32   `json:"number"`            // 挂单量
	TransactedVolume float32   `json:"transacted_volume"` // 已成交额
	TransactedAmount float32   `json:"transacted_amount"` // 已成交量
	Fee              float32   `json:"fee"`               // 手续费
	Way              int       `json:"way"`               // 买入单还是卖出单
	UserId           int       `json:"user_id"`           // 用户id
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
