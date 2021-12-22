package repository

import "github.com/anupamgogoi/go-starwars/pkg/models"

type ItemRepository interface {
	Find(id int) models.Item
	FindAll() []models.Item
	FindByInventoryId(id int) []models.Item
	DeleteItem(id int)
	UpdateItem(item models.Item)
	Save(inventory models.Item) int
}
