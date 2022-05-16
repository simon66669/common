package models

import "database/sql"

//SellerCurrencies 商家币种表
type SellerCurrencies struct {
	Id                   int          `json:"id"`
	SellerId             string       `json:"seller_id"`              // 商家id
	CurrencyId           string       `json:"currency_id"`            // 币种id
	CompletedBuy         float32      `json:"completed_buy"`          // 累计买入量
	CompletedSell        float32      `json:"completed_sell"`         // 累计卖出量
	OriginCautionBalance float32      `json:"origin_caution_balance"` // 原始保证金额度
	CautionBalance       float32      `json:"caution_balance"`        // 当前剩余保证金额度
	BuyStatus            string       `json:"buy_status"`             // 挂买权限:0.未启用,1.启用
	SellStatus           int          `json:"sell_status"`            // 挂卖权限:0.未启用,1.启用
	CreatedAt            sql.NullTime `json:"created_at"`
	UpdatedAt            sql.NullTime `json:"updated_at"`
}
