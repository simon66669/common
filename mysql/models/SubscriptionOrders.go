package models

import "database/sql"

type SubscriptionOrders struct {
	Id                int             `json:"id"`
	UserId            sql.NullInt64   `json:"user_id"`
	SubId             sql.NullInt64   `json:"sub_id"`
	Number            sql.NullFloat64 `json:"number"`          // 申购数量
	PayCurrencyId     sql.NullInt64   `json:"pay_currency_id"` // 支付币id
	PayAmount         sql.NullFloat64 `json:"pay_amount"`      // 申购金额
	PayLastPrice      sql.NullFloat64 `json:"pay_last_price"`  // 支付币最后计算价格
	CreatedAt         sql.NullTime    `json:"created_at"`
	UpdatedAt         sql.NullTime    `json:"updated_at"`
	Status            sql.NullInt64   `json:"status"`              // 1.待抽签 2.未中签 3.已中签 4.已发币 5.已退款
	WinningLotsNumber sql.NullFloat64 `json:"winning_lots_number"` // 中签数量
	IsReturn          sql.NullInt64   `json:"is_return"`           // 0.否 1.已退款 2.部分退款 3.全额退款
}
