package models

import (
	"database/sql"
	"time"
)

//TxHashes 充币哈希表
type TxHashes struct {
	Id                 int             `json:"id"`
	CurrencyId         int             `json:"currency_id"` // 币种ID
	Hash               string          `json:"hash"`        // 交易哈希
	Type               int             `json:"type"`        // 0，充币，1归拢， 2提币，3 打入手续费
	Status             int             `json:"status"`      // 0，等待中，1完成，2失败
	Memo               string          `json:"memo"`        // 备注
	UserId             int             `json:"user_id"`
	CurrencyProtocolId int             `json:"currency_protocol_id"` // 币种协议ID
	CallbackHandle     string          `json:"callback_handle"`      // 更新哈希时的回调方法
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
	Amount             sql.NullFloat64 `json:"amount"` // 数量
}
