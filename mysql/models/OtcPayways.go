package models

import "database/sql"

//OtcPayways 法币收款信息，关联detail
type OtcPayways struct {
	Id        int          `json:"id"`
	DetailId  string       `json:"detail_id"`
	UserId    string       `json:"user_id"`   // 卖方用户id
	PaywayId  string       `json:"payway_id"` // 支付方式
	RawData   string       `json:"raw_data"`  // 收款信息
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
