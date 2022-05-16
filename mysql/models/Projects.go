package models

import "database/sql"

type Projects struct {
	Id                      string          `json:"id"`
	CurrencyId              sql.NullInt64   `json:"currency_id"`              // 币种id
	MaxBuyNumber            sql.NullFloat64 `json:"max_buy_number"`           // 最大购买数量
	MinBuyNumber            sql.NullFloat64 `json:"min_buy_number"`           // 最小购买数量
	StartTime               sql.NullTime    `json:"start_time"`               // 认筹开始时间
	EndTime                 sql.NullTime    `json:"end_time"`                 // 认筹结束时间
	Banner                  sql.NullString  `json:"banner"`                   // 项目banner
	Title                   sql.NullString  `json:"title"`                    // 项目标题
	Introduction            sql.NullString  `json:"introduction"`             // 项目简介
	DetailIntroduction      sql.NullString  `json:"detail_introduction"`      // 项目详细介绍
	ParticipationConditions sql.NullString  `json:"participation_conditions"` // 参与条件
	CreatedAt               sql.NullTime    `json:"created_at"`
	UpdatedAt               sql.NullTime    `json:"updated_at"`
	ProjectNumber           sql.NullFloat64 `json:"project_number"`     // 总认购数量
	QuantityPurchased       sql.NullFloat64 `json:"quantity_purchased"` // 已购数量
}
