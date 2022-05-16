package models

type AdminRoute struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	IsAddLog int    `json:"is_add_log"`
}
