package models

//AdminRoleModules 角色访问权限表
type AdminRoleModules struct {
	Id          string `json:"id"`
	AdminRoleId int    `json:"admin_role_id"`
	ModuleId    int    `json:"module_id"`
}
