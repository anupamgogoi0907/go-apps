package db

import (
	"database/sql"
	"github.com/anupamgogoi/go-starwars/pkg/models"
	"log"
)

type MySQLStorageItem struct {
	DB *sql.DB
}

func (s MySQLStorageItem) Find(id int) models.Item {
	item := models.Item{}
	return item
}
func (s MySQLStorageItem) FindAll() []models.Item {
	var items []models.Item
	query := "SELECT * FROM Item"
	row, err := s.DB.Query(query)
	if err != nil {
		log.Panicln(err)
	}
	for row.Next() {
		var id int
		var name string
		var quantity int
		var points int
		err = row.Scan(&id, &name, &quantity, &points)

		itemRes := models.Item{Id: id,
			Name:   name,
			Qty:    quantity,
			Points: points,
		}
		items = append(items, itemRes)

	}
	return items
}
func (s MySQLStorageItem) FindByInventoryId(id int) []models.Item {
	return nil
}
func (s MySQLStorageItem) DeleteItem(id int) {

}
func (s MySQLStorageItem) UpdateItem(item models.Item) {

}

func (s MySQLStorageItem) Save(item models.Item) int {
	log.Default().Println("Creating a Inventory")
	sql := "INSERT INTO Item (id,name,quantity,points,inventory_id) VALUES (?,?,?,?,?)"
	stmt, err := s.DB.Prepare(sql)
	if err != nil {
		log.Panicln(err)
	}
	_, err = stmt.Exec(item.Id, item.Name, item.Qty, item.Points, item.Inventory_id)
	if err != nil {
		log.Panicln(err)
	}
	return item.Id
}
