package models

import "time"

//TxCompletes 全站交易表
type TxCompletes struct {
	Id              int       `json:"id"`
	CurrencyMatchId int       `json:"currency_match_id"` // 交易对
	Price           float32   `json:"price"`             // 价格
	Number          float32   `json:"number"`            // 交易量
	OutUserId       int       `json:"out_user_id"`       // 卖出方
	InUserId        int       `json:"in_user_id"`        // 买入方
	Way             int       `json:"way"`               // 买入撮合还是卖出撮合
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
