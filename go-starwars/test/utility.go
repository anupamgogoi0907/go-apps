package test

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
)

func GetConnection() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Panicln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Panicln(err)
	}
	// Create tables
	InitDatabase(db)
	return db
}

func InitDatabase(db *sql.DB) {
	dir, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}
	log.Println(dir)
	script, err := ioutil.ReadFile(dir + "/resources/sqlite_create_tables.sql")
	if err != nil {
		log.Panicln(err)
	}
	query := string(script)

	_, errQuery := db.Exec(query)
	if errQuery != nil {
		log.Panicln(errQuery)
	}
}
