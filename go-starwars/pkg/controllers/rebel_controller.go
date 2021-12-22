package controllers

import (
	"database/sql"
	"github.com/anupamgogoi/go-starwars/pkg/models"
	"github.com/anupamgogoi/go-starwars/pkg/repository/db"
	"github.com/gorilla/mux"
	"net/http"
)

type RebelController struct {
	Router *mux.Router
	DB     *sql.DB
}

func (controller RebelController) GetRebels(w http.ResponseWriter, r *http.Request) {
	repo := db.MySQLStorageRebel{
		DB: controller.DB,
	}
	repo.FindAll()
}

func (controller RebelController) AddRebel(w http.ResponseWriter, r *http.Request) {
	repo := db.MySQLStorageRebel{
		DB: controller.DB,
	}
	repo.Save(models.Rebel{})
}

func (controller RebelController) InitializeRoutes() {
	controller.Router.HandleFunc("/rebel", controller.GetRebels).Methods("GET")
	controller.Router.HandleFunc("/rebel", controller.AddRebel).Methods("POST")
}
