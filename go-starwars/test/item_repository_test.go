package test

import (
	"github.com/anupamgogoi/go-starwars/pkg/models"
	repo "github.com/anupamgogoi/go-starwars/pkg/repository/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItemSave(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	s := repo.MySQLStorageItem{DB: db}
	item := models.Item{
		Id:           1,
		Name:         "Arma",
		Qty:          1,
		Points:       4,
		Inventory_id: 1,
	}
	id := s.Save(item)
	assert.Equal(t, item.Id, id)
}
