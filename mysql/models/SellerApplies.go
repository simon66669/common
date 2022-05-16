package models

import "database/sql"

//SellerApplies 商家申请记录
type SellerApplies struct {
	Id         int          `json:"id"`
	UserId     string       `json:"user_id"`    // 用户id
	Name       string       `json:"name"`       // 名称
	Currencies string       `json:"currencies"` // 币种id
	Status     int          `json:"status"`     // 状态:0.已提交,1.通过,2.驳回
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}
