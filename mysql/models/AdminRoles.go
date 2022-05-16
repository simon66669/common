package models

//AdminRoles 系统角色表
type AdminRoles struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	IsSuper string `json:"is_super"`
}
