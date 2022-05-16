package models

import "time"

//NotifyLogs 发送通知记录
type NotifyLogs struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`    // 内容
	To        string    `json:"to"`         // 发给谁的
	Type      int       `json:"type"`       // 类型
	ExtraData string    `json:"extra_data"` // 预留数据
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
