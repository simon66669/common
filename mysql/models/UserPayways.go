package models

import "database/sql"

//UserPayways 用户支付/收款方式
type UserPayways struct {
	Id        int          `json:"id"`
	UserId    string       `json:"user_id"`   // 用户id
	PaywayId  string       `json:"payway_id"` // 支付方式
	FormId    string       `json:"form_id"`
	RawData   string       `json:"raw_data"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
