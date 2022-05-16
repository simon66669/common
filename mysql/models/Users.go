package models

import (
	"database/sql"
	"time"
)

//Users 用户表
type Users struct {
	Id                         int            `json:"id"`
	Uid                        string         `json:"uid"`           // uid
	Mobile                     string         `json:"mobile"`        // 手机
	Email                      string         `json:"email"`         // 邮箱
	AreaId                     int            `json:"area_id"`       // 区域id
	Password                   string         `json:"password"`      // 密码
	PayPassword                string         `json:"pay_password"`  // 支付密码
	Status                     int            `json:"status"`        // 用户状态1正常2锁定
	TxStatus                   int            `json:"tx_status"`     // 是否开启交易
	ParentId                   int            `json:"parent_id"`     // 上级id
	InviteCode                 string         `json:"invite_code"`   // 邀请码
	LastLoginIp                string         `json:"last_login_ip"` // 最后登录ip
	LastLoginAt                time.Time      `json:"last_login_at"` // 最后登录时间
	ParentsPath                string         `json:"parents_path"`  // 用户父级关系
	AgentId                    int            `json:"agent_id"`      // 对应代理商id
	AgentNodeId                int            `json:"agent_node_id"` // 所属上级代理商id
	AgentPath                  string         `json:"agent_path"`    // 上级代理商关系
	UpdatedAt                  time.Time      `json:"updated_at"`
	CreatedAt                  time.Time      `json:"created_at"`
	IsCreateAddress            sql.NullInt64  `json:"is_create_address"`             // 是否创建了地址
	MicroContStatus            sql.NullInt64  `json:"micro_cont_status"`             // 是否单控概率
	MicroRiskProfitProbability sql.NullInt64  `json:"micro_risk_profit_probability"` // 概率控盈利概率%
	CreditScore                sql.NullInt64  `json:"credit_score"`                  // 信用分
	Level                      sql.NullString `json:"level"`                         // 等级
	HeadImage                  sql.NullString `json:"head_image"`                    // 头像
}
