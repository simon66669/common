package models

import "database/sql"

//ChainProtocols 链上协议
type ChainProtocols struct {
	Id           int            `json:"id"`
	Code         string         `json:"code"`      // 协议代码
	Classname    string         `json:"classname"` // 驱动
	Data         string         `json:"data"`      // 拓展信息,比如说手续费
	MainCoinType sql.NullString `json:"main_coin_type"`
	CoinType     sql.NullString `json:"coin_type"`
}
