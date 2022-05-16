package models

import (
	"database/sql"
	"time"
)

//Currencies 系统币种表
type Currencies struct {
	Id              int             `json:"id"`
	Code            string          `json:"code"`             // 币种符号
	Logo            string          `json:"logo"`             // logo
	Desc            string          `json:"desc"`             // 简介
	Decimal         int             `json:"decimal"`          // 币种小数位
	AccountsDisplay string          `json:"accounts_display"` // 都展示什么账户
	IsQuote         int             `json:"is_quote"`         // 是否是计价币
	UsdPrice        float32         `json:"usd_price"`        // 币种美元价格
	CnyPrice        float32         `json:"cny_price"`        // 币种人民币价格
	DrawRate        float32         `json:"draw_rate"`        // 提币费率
	DrawMin         float32         `json:"draw_min"`         // 最小提币量
	DrawMax         float32         `json:"draw_max"`         // 最大提币量
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	OpenDraw        int             `json:"open_draw"`        // 是否开启提币
	MicroTradeFee   float32         `json:"micro_trade_fee"`  // 秒合约下单费率
	OpenRecharge    int             `json:"open_recharge"`    // 是否开启充币
	Sort            int             `json:"sort"`             // 排序
	IsPb            sql.NullInt64   `json:"is_pb"`            // 是平台币
	MainCoinType    sql.NullString  `json:"main_coin_type"`   // 主币种编号
	CoinType        sql.NullString  `json:"coin_type"`        // 子币种编号
	IsDb            sql.NullInt64   `json:"is_db"`            // 0不是平台币 1是
	OptionTradeFee  sql.NullFloat64 `json:"option_trade_fee"` // 期权合约下单费率
	IsNew           sql.NullInt64   `json:"is_new"`           // 0不是新发币 1是
}
