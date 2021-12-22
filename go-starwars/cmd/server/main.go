package main

import (
	"database/sql"
	"github.com/anupamgogoi/go-starwars/pkg/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/starwars_go")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	router := mux.NewRouter()
	conn := ConnectDB()
	rebelController := controllers.RebelController{
		Router: router,
		DB:     conn,
	}
	rebelController.InitializeRoutes()

}
