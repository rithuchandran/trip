package domain

type Trip struct {
	Id          int    `json:"id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}
