package models

import "database/sql"

//OtcActions 法币交易各环节交易快照
type OtcActions struct {
	Id           int          `json:"id"`
	DetailId     string       `json:"detail_id"`     // 交易id
	UserId       string       `json:"user_id"`       // 触发操作的用户id
	BuyUserId    string       `json:"buy_user_id"`   // 买方id
	SellUserId   string       `json:"sell_user_id"`  // 卖方id
	Status       string       `json:"status"`        // 参见模型定义
	OperatorType string       `json:"operator_type"` // 操作员类型:0.用户,1.系统,2.后台管理员
	PublicMsg    string       `json:"public_msg"`    // 买卖双方共享的消息
	ToBuyMsg     string       `json:"to_buy_msg"`    // 向卖方发送的消息
	ToSellMsg    string       `json:"to_sell_msg"`   // 向买方发送的消息
	Memo         string       `json:"memo"`          // 说明
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}
