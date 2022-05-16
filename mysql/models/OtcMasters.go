package models

import "database/sql"

//OtcMasters 法币挂单广告
type OtcMasters struct {
	Id           int            `json:"id"`
	Way          sql.NullString `json:"way"` // 发布方向
	UserId       string         `json:"user_id"`
	SellerId     string         `json:"seller_id"`     // 商家id
	CurrencyId   string         `json:"currency_id"`   // 发布币种
	AreaId       string         `json:"area_id"`       // 商家区域,限定交易的实际兑现币种符号
	Price        float32        `json:"price"`         // 价格
	TotalQty     float32        `json:"total_qty"`     // 总数量
	CompletedQty float32        `json:"completed_qty"` // 已完成数量
	DealingQty   float32        `json:"dealing_qty"`   // 正在交易数量
	RemainQty    float32        `json:"remain_qty"`    // 剩余数量
	MinNumber    float32        `json:"min_number"`    // 最小交易量
	MaxNumber    float32        `json:"max_number"`    // 最大交易量
	Status       int            `json:"status"`        // 状态:0.暂停交易,1.交易中,2.下架(被下架管理员才能恢复),3.完成,4.取消
	Payways      string         `json:"payways"`       // 支持的付款方式
	CreatedAt    sql.NullTime   `json:"created_at"`    // 发布时间
	StartedAt    sql.NullTime   `json:"started_at"`    // 开始交易时间
	PausedAt     sql.NullTime   `json:"paused_at"`     // 暂停交易时间
	StopedAt     sql.NullTime   `json:"stoped_at"`     // 下架交易时间
	CanceledAt   sql.NullTime   `json:"canceled_at"`   // 撤单时间
	FinishedAt   sql.NullTime   `json:"finished_at"`   // 完成时间
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}
