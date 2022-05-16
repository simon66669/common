package models

import "database/sql"

type SetCommissions struct {
	Id        string         `json:"id"`
	Levels    sql.NullString `json:"levels"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	Status    sql.NullInt64  `json:"status"`
}
