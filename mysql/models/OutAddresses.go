package models

import "database/sql"

type OutAddresses struct {
	Id                 int             `json:"id"`
	UserId             sql.NullInt64   `json:"user_id"`              // 用户
	CurrencyId         sql.NullInt64   `json:"currency_id"`          // 币种id
	CurrencyProtocolId sql.NullInt64   `json:"currency_protocol_id"` // 区块链协议id
	OutAddress         sql.NullString  `json:"out_address"`          // 提币地址
	Fee                sql.NullFloat64 `json:"fee"`                  // 手续费
	Meno               sql.NullString  `json:"meno"`                 // 备注
	CreatedAt          sql.NullTime    `json:"created_at"`
	UpdatedAt          sql.NullTime    `json:"updated_at"`
}
