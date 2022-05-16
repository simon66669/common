package models

import "database/sql"

type AgentMoneyLogs struct {
	Id              int          `json:"id"`
	AgentId         int          `json:"agent_id"`          // 所属代理商
	Type            int          `json:"type"`              // 类型。1.代理商头寸，2代理商手续费
	RelateId        string       `json:"relate_id"`         // 关联id。比如订单id等
	Change          float32      `json:"change"`            // 本次变动
	Memo            string       `json:"memo"`              // 备注
	UserId          int          `json:"user_id"`           // 贡献收益的用户id
	Status          int          `json:"status"`            // 0  是否提现到账
	QuoteCurrencyId int          `json:"quote_currency_id"` // 法币id
	CreatedAt       sql.NullTime `json:"created_at"`        // 结算时间
	UpdatedAt       sql.NullTime `json:"updated_at"`        // 提现到账时间
}
