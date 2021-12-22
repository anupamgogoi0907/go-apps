package test

import (
	"github.com/anupamgogoi/go-starwars/pkg/models"
	repo "github.com/anupamgogoi/go-starwars/pkg/repository/db"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRebelSave(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	storage := repo.MySQLStorageRebel{
		DB: db,
	}
	rebel := models.Rebel{
		Id:              1,
		Name:            "Anupam",
		Gender:          "M",
		Age:             33,
		Traitor:         false,
		Localization_id: 1,
		Inventory_id:    1,
	}
	storage.Save(rebel)
}

func TestRebelFindAll(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	storage := repo.MySQLStorageRebel{
		DB: db,
	}
	rebel := models.Rebel{
		Id:              1,
		Name:            "Anupam",
		Gender:          "M",
		Age:             33,
		Traitor:         false,
		Localization_id: 1,
		Inventory_id:    1,
	}
	storage.Save(rebel)

	rebels := storage.FindAll()
	assert.NotEmpty(t, len(rebels))
}
