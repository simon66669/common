package models

import "time"

//FailedJobs 队列失败任务表
type FailedJobs struct {
	Id         string    `json:"id"`
	Connection string    `json:"connection"`
	Queue      string    `json:"queue"`
	Payload    string    `json:"payload"`
	Exception  string    `json:"exception"`
	FailedAt   time.Time `json:"failed_at"`
}
