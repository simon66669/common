package models

import "database/sql"

type SellerCurrencyApplies struct {
	Id         int          `json:"id"`
	UserId     string       `json:"user_id"`    // 用户id
	SellerId   string       `json:"seller_id"`  // 商家id
	Currencies string       `json:"currencies"` // 申请的币种列表
	Status     int          `json:"status"`     // 状态:0.申请中,1.通过,2.驳回
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}
