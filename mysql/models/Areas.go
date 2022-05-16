package models

import "database/sql"

//Areas 区域信息
type Areas struct {
	Id             int          `json:"id"`
	Code           string       `json:"code"`            // 区域代码
	Name           string       `json:"name"`            // 区域名称
	Currency       string       `json:"currency"`        // 区域货币代码
	CurrencySymbol string       `json:"currency_symbol"` // 区域货币符号
	Decimal        int          `json:"decimal"`         // 货币小数位数
	GlobalCode     string       `json:"global_code"`     // 国际电信区号
	LangId         string       `json:"lang_id"`         // 语言id
	Payways        string       `json:"payways"`         // 支持的支付方式
	Timezone       string       `json:"timezone"`        // 时区信息
	Sort           int          `json:"sort"`
	CreatedAt      sql.NullTime `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
}
