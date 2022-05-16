package models

import (
	"database/sql"
	"time"
)

//AccountDraws 提币记录表
type AccountDraws struct {
	Id                 int            `json:"id"`
	CurrencyId         int            `json:"currency_id"` // 币种ID
	UserId             int            `json:"user_id"`     // 用户ID
	Address            string         `json:"address"`     // 提币地址
	Number             float32        `json:"number"`      // 提币数量
	Status             int            `json:"status"`      // 状态1待审核2已通过3已驳回
	Fee                float32        `json:"fee"`         // 手续费
	RealNumber         float32        `json:"real_number"` // 真实到账数量
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	ReviewAt           int            `json:"review_at"`     // 审核操作时间
	TxHashId           int            `json:"tx_hash_id"`    // 关联hash表id
	UseChainApi        int            `json:"use_chain_api"` // 是否走链上   0否  1是
	Memo               string         `json:"memo"`
	Notes              string         `json:"notes"`                // 说明
	AccountTypeId      int            `json:"account_type_id"`      // 账户类型ID
	CurrencyProtocolId int            `json:"currency_protocol_id"` // 币种协议ID
	FinishTime         sql.NullTime   `json:"finish_time"`
	BusinessId         sql.NullString `json:"businessId"`
}
