package models

import "database/sql"

type ProjectRobotPeriods struct {
	Id            string          `json:"id"`
	DayOpenPrice  sql.NullFloat64 `json:"day_open_price"`
	Period        sql.NullString  `json:"period"`
	UpdatedAt     sql.NullTime    `json:"updated_at"`
	CreatedAt     sql.NullTime    `json:"created_at"`
	RobotId       sql.NullInt64   `json:"robot_id"`
	Change        sql.NullFloat64 `json:"change"`
	DayClosePrice sql.NullFloat64 `json:"day_close_price"`
}
