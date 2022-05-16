package models

import "time"

//TxOuts 交易卖出表
type TxOuts struct {
	Id               int       `json:"id"`
	CurrencyMatchId  int       `json:"currency_match_id"` // 交易对id
	UserId           int       `json:"user_id"`           // 用户id
	Price            float32   `json:"price"`             // 价格
	Number           float32   `json:"number"`
	TransactedVolume float32   `json:"transacted_volume"` // 已成交金量
	TransactedAmount float32   `json:"transacted_amount"` // 已成交额
	QuoteAccountId   int       `json:"quote_account_id"`
	BaseAccountId    int       `json:"base_account_id"` // 钱包id
	Fee              float32   `json:"fee"`             // 手续费
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
