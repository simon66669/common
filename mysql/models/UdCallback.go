package models

import "database/sql"

type UdCallback struct {
	Id           int            `json:"id"`
	Address      sql.NullString `json:"address"`
	Amount       sql.NullString `json:"amount"`
	Decimals     sql.NullString `json:"decimals"`
	Fee          sql.NullString `json:"fee"`
	CoinType     sql.NullString `json:"coinType"`
	MainCoinType sql.NullString `json:"mainCoinType"`
	BusinessId   sql.NullString `json:"businessId"`
	BlockHigh    sql.NullString `json:"blockHigh"`
	TradeId      sql.NullString `json:"tradeId"`
	TradeType    sql.NullInt64  `json:"tradeType"`
	Txid         sql.NullString `json:"txid"`
	Status       sql.NullInt64  `json:"status"`
	Memo         sql.NullString `json:"memo"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
	UserId       sql.NullInt64  `json:"user_id"`
}
