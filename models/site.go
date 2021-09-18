package models

type Site struct {
	ID        int64  `json:"ID"`
	SiteName  string `json:"site_name"`
	SiteEmail string `json:"site_email"`
	Slogan    string `json:"slogan"`
	Notice    string `json:"notice"`
}
