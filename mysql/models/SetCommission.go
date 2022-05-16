package models

import "database/sql"

type SetCommission struct {
	Id           string          `json:"id"`
	CreatedAt    sql.NullTime    `json:"created_at"`
	UpdatedAt    sql.NullTime    `json:"updated_at"`
	CustomerRate sql.NullFloat64 `json:"customer_rate"`
	AgentRate    sql.NullFloat64 `json:"agent_rate"`
}
