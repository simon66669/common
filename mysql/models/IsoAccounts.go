package models

import "time"

//IsoAccounts 账户表
type IsoAccounts struct {
	Id            int       `json:"id"`
	UserId        int       `json:"user_id"`
	CurrencyId    int       `json:"currency_id"`     // 币种id
	AccountTypeId int       `json:"account_type_id"` // 账户类型id
	Balance       float32   `json:"balance"`
	LockBalance   float32   `json:"lock_balance"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
