package models

//CurrencyQuotations 币种涨跌行情表
type CurrencyQuotations struct {
	Id              string  `json:"id"`
	CurrencyMatchId int     `json:"currency_match_id"` // 交易对
	Change          float32 `json:"change"`            // 涨跌幅
	Volume          float32 `json:"volume"`            // 成交量
	Close           float32 `json:"close"`             // 当前价位
	Open            float32 `json:"open"`              // 开盘价
	Amount          float32 `json:"amount"`            // 成交金额
	High            float32 `json:"high"`              // 最高价
	Low             float32 `json:"low"`               // 最低价
}
