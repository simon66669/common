package models

//Langs 语言表
type Langs struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status int    `json:"status"` // 1，启动，0禁用
	Sort   int    `json:"sort"`   // 排序
}
