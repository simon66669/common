package models

import "database/sql"

//Sellers 法币商家
type Sellers struct {
	Id        int          `json:"id"`
	UserId    string       `json:"user_id"` // 用户id
	Name      string       `json:"name"`    // 名称
	Status    int          `json:"status"`  // 状态:0.未启用,1.启用
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
