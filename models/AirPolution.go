package models

type AirPolution struct {
	Id        int    `json:"id"`
	Lat       string `json:"lat"`
	Lng       string `json:"lng"`
	Compound  string `json:"compound"`
	CreatedAt string `json:"created_at"`
}
