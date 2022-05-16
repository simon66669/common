package models

import "time"

//ChainWallets 钱包表
type ChainWallets struct {
	Id                 int       `json:"id"`
	UserId             int       `json:"user_id"`              // 币种id
	CurrencyId         int       `json:"currency_id"`          // 币种id
	CurrencyProtocolId int       `json:"currency_protocol_id"` // 币种链上协议id
	ChainProtocolId    int       `json:"chain_protocol_id"`    // 链上协议id
	Address            string    `json:"address"`
	PrivateKey         string    `json:"private_key"` // 私钥
	Balance            float32   `json:"balance"`     // 链上余额
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Memo               string    `json:"memo"` // 可用于EOS账户充币
}
