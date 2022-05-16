package models

import "time"

//FeedbackTypes 反馈类别表
type FeedbackTypes struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`  // 类别名称
	Sorts     int       `json:"sorts"` // 排序
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
