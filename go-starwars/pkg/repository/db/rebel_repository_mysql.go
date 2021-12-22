package db

import (
	"database/sql"
	"github.com/anupamgogoi/go-starwars/pkg/models"
	"github.com/anupamgogoi/go-starwars/pkg/repository"
	"log"
)

type MySQLStorageRebel struct {
	DB *sql.DB
}

func (s MySQLStorageRebel) FindAll() []models.Rebel {
	log.Default().Println("Finding all Rebels")

	var rebels []models.Rebel

	sql := "SELECT * FROM Rebel"
	stmt, err := s.DB.Prepare(sql)
	if err != nil {
		log.Panicln(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		if err != nil {
			log.Panicln(err)
		}
	}
	for rows.Next() {
		var Id int
		var Name string
		var Age int
		var Gender string
		var Traitor bool
		var Localization_id int
		var Inventory_id int
		err = rows.Scan(&Id, &Name, &Age, &Gender, &Traitor, &Localization_id, &Inventory_id)
		if err != nil {
			log.Panicln(err)
		}
		rebel := models.Rebel{
			Id:              Id,
			Name:            Name,
			Age:             Age,
			Gender:          Gender,
			Traitor:         Traitor,
			Localization_id: Localization_id,
			Inventory_id:    Inventory_id,
		}
		rebels = append(rebels, rebel)
	}
	return rebels
}
func (s MySQLStorageRebel) FindById(it int) models.Rebel {
	return models.Rebel{}
}

func (s MySQLStorageRebel) Save(rebel models.Rebel) int {
	log.Default().Println("Creating a Rebel")
	sql := "INSERT INTO Rebel (id,name,age,gender,traitor,localization_id,inventory_id) VALUES (?,?,?,?,?,?,?)"
	stmt, err := s.DB.Prepare(sql)
	if err != nil {
		log.Panicln(err)
	}
	_, err = stmt.Exec(rebel.Id, rebel.Name, rebel.Age, rebel.Gender, rebel.Traitor, rebel.Localization_id, rebel.Inventory_id)
	if err != nil {
		log.Panicln(err)
	}
	return rebel.Id
}
func (s MySQLStorageRebel) Delete(id int) {

}
func (s MySQLStorageRebel) Update(rebel models.Rebel) {

}
func (s MySQLStorageRebel) UpdateReporters(id int, reporters string) {

}
func (s MySQLStorageRebel) UpdateTraitor(id int, traitor bool) {

}

func Test1() {
	var test repository.RebelRepository
	test = MySQLStorageRebel{}
	test.Delete(1)
}
