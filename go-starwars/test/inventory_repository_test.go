package test

import (
	"github.com/anupamgogoi/go-starwars/pkg/models"
	repo "github.com/anupamgogoi/go-starwars/pkg/repository/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInventorySave(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	inv := models.Inventory{
		Id: 1,
	}
	s := repo.MySQLStorageInventory{
		DB: db,
	}
	id := s.Save(inv)
	assert.Equal(t, 1, id, "Passsssssss")
}

func TestInventoryFindById(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	inv := models.Inventory{
		Id: 1,
	}
	s := repo.MySQLStorageInventory{
		DB: db,
	}
	id := s.Save(inv)
	result := s.FindById(id)
	assert.Equal(t, inv.Id, result.Id)
}
