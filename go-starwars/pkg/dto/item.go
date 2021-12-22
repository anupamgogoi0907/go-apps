package dto

type Item struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Qty    int    `json:"qty"`
	Points int    `json:"points"`
}
