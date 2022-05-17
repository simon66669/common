package models

import (
	"database/sql"
	"time"
)

//Robots 机器人表
type Robots struct {
	Id              int             `json:"id"`
	CurrencyMatchId int             `json:"currency_match_id"`
	Valume          float32         `json:"valume"` // 单次交易量
	UserId          int             `json:"user_id"`
	Price           float32         `json:"price"`          // 预期价格
	VirtualSymbol   string          `json:"virtual_symbol"` // 虚拟交易对
	Match           int             `json:"match"`          // 是否撮合
	ProcessId       int             `json:"process_id"`     // 机器人使用的进程id
	Status          int             `json:"status"`
	Decimal         int             `json:"decimal"` // 价格的小数位数
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	Point           sql.NullFloat64 `json:"point"`
	NarrowMultiple  sql.NullFloat64 `json:"narrow_multiple"` // 缩小倍数

	CurrencyMatch CurrencyMatches `gorm:"foreignKey:currency_match_id"`
}
