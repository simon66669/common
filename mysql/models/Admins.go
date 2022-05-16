package models

import "time"

//Admins 系统用户表
type Admins struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	RoleId       int       `json:"role_id"`
	LastLoginAt  time.Time `json:"last_login_at"`
	LastLoginIp  string    `json:"last_login_ip"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
	GoogleSecret string    `json:"google_secret"` // 谷歌验证器,密码
}
