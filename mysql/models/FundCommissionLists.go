package models

import "database/sql"

type FundCommissionLists struct {
	Id                string          `json:"id"`
	ForSubId          sql.NullInt64   `json:"for_sub_id"`
	ForUserId         sql.NullInt64   `json:"for_user_id"`
	ForPeriodsId      sql.NullInt64   `json:"for_periods_id"`
	UserId            sql.NullInt64   `json:"user_id"`
	ForFundId         sql.NullInt64   `json:"for_fund_id"`
	CreatedAt         sql.NullTime    `json:"created_at"`
	Amount            sql.NullFloat64 `json:"amount"`
	CommissionLevel   sql.NullInt64   `json:"commission_level"`
	CommissionSetData sql.NullString  `json:"commission_set_data"`
}
