package models

//Settings 系统设置表
type Settings struct {
	Id            string `json:"id"`
	Key           string `json:"key"`
	Value         string `json:"value"`
	Desc          string `json:"desc"`
	SettingTypeId int    `json:"setting_type_id"`
	Visible       int    `json:"visible"` // 前台是否可见
}
