package dto

type Rebel struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	Traitor   bool      `json:"traitor"`
	Local     Local     `json:"local"`
	Inventory Inventory `json:"inventory_id"`
}
