package search

import (
	"CarRental/entities"
)

type SearchService interface {
	GetByName(keyword string) ([]entities.Car, error)
}