package db

import (
	"database/sql"
	"github.com/anupamgogoi/go-starwars/pkg/models"
	"github.com/anupamgogoi/go-starwars/pkg/repository"
	"log"
)

type MySQLStorageInventory struct {
	DB *sql.DB
}

func (s MySQLStorageInventory) FindById(id int) models.Inventory {
	log.Default().Println("Creating a Inventory")
	query := "SELECT * FROM Inventory WHERE id=?"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		log.Panicln(err)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		log.Panicln(err)
	}
	result := models.Inventory{}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Panicln(err)
		}
		result.Id = id

	}
	return result
}

func (s MySQLStorageInventory) Save(inventory models.Inventory) int {
	log.Default().Println("Creating a Inventory")
	sql := "INSERT INTO Inventory (id) VALUES (?)"
	stmt, err := s.DB.Prepare(sql)
	if err != nil {
		log.Panicln(err)
	}
	_, err = stmt.Exec(inventory.Id)
	if err != nil {
		log.Panicln(err)
	}
	return inventory.Id
}
func (s MySQLStorageInventory) Update(inventory models.Inventory) {

}
func (s MySQLStorageInventory) Delete(id int) {

}

func Test() {
	var test repository.InventoryRepository
	test = MySQLStorageInventory{}
	test.Delete(1)
}
