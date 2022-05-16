package models

import "database/sql"

//OtcDetails 法币交易记录
type OtcDetails struct {
	Id           int            `json:"id"`
	MasterId     string         `json:"master_id"`      // 商家订单id
	CurrencyId   string         `json:"currency_id"`    // 币种id
	Way          sql.NullString `json:"way"`            // 用户的交易方向,与商家的订单方向相反
	UserId       string         `json:"user_id"`        // 用户id
	SellerId     string         `json:"seller_id"`      // 商家id
	SellerUserId string         `json:"seller_user_id"` // 商家的用户id
	BuyUserId    string         `json:"buy_user_id"`    // 买方id
	SellUserId   string         `json:"sell_user_id"`   // 卖方id(不一定是商家)
	Price        float32        `json:"price"`          // 价格
	Number       float32        `json:"number"`         // 数量
	Amount       float32        `json:"amount"`         // 金额
	PayVoucher   string         `json:"pay_voucher"`    // 支付凭证
	Status       int            `json:"status"`         // 状态:0.交易中,1.已付款待确认,2.超时进入延期,3.进入椎权,4.已取消,5.已确认
	CanceledAt   sql.NullTime   `json:"canceled_at"`    // 取消时间
	PayedAt      sql.NullTime   `json:"payed_at"`       // 支付时间
	DeferedAt    sql.NullTime   `json:"defered_at"`     // 延期时间
	ArbitratedAt sql.NullTime   `json:"arbitrated_at"`  // 仲裁时间
	FinishedAt   sql.NullTime   `json:"finished_at"`    // 完成时间
	CreatedAt    sql.NullTime   `json:"created_at"`     // 创建时间
	UpdatedAt    sql.NullTime   `json:"updated_at"`     // 更新时间
}
