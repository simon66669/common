package models

type CurrencyProtocols struct {
	Id              int    `json:"id"`
	ChainProtocolId int    `json:"chain_protocol_id"`
	CurrencyId      int    `json:"currency_id"`
	InAddress       string `json:"in_address"`      // 入金账号
	OutAddress      string `json:"out_address"`     // 出金账号
	OutPrivateKey   string `json:"out_private_key"` // 出金账号私钥
	Decimal         int    `json:"decimal"`         // 小数位数
	TokenAddress    string `json:"token_address"`
}
