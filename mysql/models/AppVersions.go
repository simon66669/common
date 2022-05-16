package models

import "time"

//AppVersions APP版本管理
type AppVersions struct {
	Id            int       `json:"id"`
	VersionNumber string    `json:"version_number"`
	PkgUrl        string    `json:"pkg_url"`
	WgtUrl        string    `json:"wgt_url"`
	DownloadUrl   string    `json:"download_url"`
	Type          int       `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DownloadType  string    `json:"download_type"`
	OtherUrl      string    `json:"other_url"`
}
