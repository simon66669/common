package models

type AccountTypes struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Classname   string `json:"classname"`
	AccountCode string `json:"account_code"`
	IsRecharge  int    `json:"is_recharge"` // 是否可用与到账 只能存在一个1到账2不到账
	IsDisplay   string `json:"is_display"`  // 是否显示
}
