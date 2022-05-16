package models

type LeverMultiple struct {
	Id    int     `json:"id"`
	Type  int     `json:"type"` // 1倍数  2手数
	Value float32 `json:"value"`
}
