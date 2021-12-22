package repository

import (
	"github.com/anupamgogoi/go-starwars/pkg/models"
)

type InventoryRepository interface {
	FindById(id int) models.Inventory
	Save(inventory models.Inventory) int
	Update(inventory models.Inventory)
	Delete(id int)
}
