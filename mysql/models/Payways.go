package models

import "database/sql"

//Payways 支付方式
type Payways struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"` // 名称
	Code      string       `json:"code"` // 支付代码
	Logo      string       `json:"logo"`
	FormId    int          `json:"form_id"` // 所需字段,关联一个表单
	FormField string       `json:"form_field"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
