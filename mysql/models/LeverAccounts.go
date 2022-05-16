package models

import (
	"database/sql"
	"time"
)

//LeverAccounts 账户表
type LeverAccounts struct {
	Id            int           `json:"id"`
	UserId        int           `json:"user_id"`
	CurrencyId    int           `json:"currency_id"` // 币种id
	Balance       float32       `json:"balance"`
	AccountTypeId int           `json:"account_type_id"` // 账户类型id
	LockBalance   float32       `json:"lock_balance"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	PShow         sql.NullInt64 `json:"p_show"`
}
