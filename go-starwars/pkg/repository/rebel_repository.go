package repository

import (
	"github.com/anupamgogoi/go-starwars/pkg/models"
)

type RebelRepository interface {
	FindAll() []models.Rebel
	FindById(id int) models.Rebel
	Save(rebel models.Rebel) int
	Delete(id int)
	Update(rebel models.Rebel)
	UpdateReporters(id int, reporters string)
	UpdateTraitor(id int, traitor bool)
}
