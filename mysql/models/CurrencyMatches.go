package models

import (
	"database/sql"
	"time"
)

//CurrencyMatches 交易对表
type CurrencyMatches struct {
	Id                    int             `json:"id"`
	QuoteCurrencyId       int             `json:"quote_currency_id"`        // 计价币种id
	BaseCurrencyId        int             `json:"base_currency_id"`         // 交易币种id
	MarketFrom            int             `json:"market_from"`              // 行情来自
	OpenLever             int             `json:"open_lever"`               // 是否开启杠杆交易
	OpenChange            int             `json:"open_change"`              // 开启币币交易
	OpenIso               int             `json:"open_iso"`                 // 开启逐仓合约
	OpenMicrotrade        int             `json:"open_microtrade"`          // 开启秒合约交易
	LeverPointRate        float32         `json:"lever_point_rate"`         // 杠杆交易点差率
	LeverOvernightFeeRate float32         `json:"lever_overnight_fee_rate"` // 杠杆交易隔夜费率
	LeverFeeRate          float32         `json:"lever_fee_rate"`           // 杠杆交易手续费率
	ChangeFeeRate         float32         `json:"change_fee_rate"`          // 币币交易手续费率
	UseProcess            int             `json:"use_process"`              // 使用的撮合队列号
	LeverMinShare         string          `json:"lever_min_share"`          // 杠杆最低手数
	LeverMaxShare         string          `json:"lever_max_share"`          // 杠杆最高手数
	LeverShareNum         int             `json:"lever_share_num"`          // 每手数量
	AutoMatch             int             `json:"auto_match"`               // 是否自动吃单,行情来自机器人或火币有效
	MatchUserId           int             `json:"match_user_id"`
	Sort                  int             `json:"sort"`            // 排序
	OrderRiskRate         float32         `json:"order_risk_rate"` // 秒合约单控浮动率
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at"`
	FloatingPoint         sql.NullFloat64 `json:"floating_point"` // 浮点
	OpenOption            sql.NullInt64   `json:"open_option"`    // 开启期权交易
	MarketTime            sql.NullTime    `json:"market_time"`
	IsNew                 sql.NullInt64   `json:"is_new"`
}
