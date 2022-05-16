package models

//AdminModuleActions 系统模块的操作
type AdminModuleActions struct {
	Id            string `json:"id"`
	AdminModuleId string `json:"admin_module_id"`
	Action        string `json:"action"`
}
