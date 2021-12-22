package dto

type Local struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}
