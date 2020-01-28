package search

import (
	"CarRental/entities"
)

type SearchRepository interface {
	GetByName(keyword string) ([]entities.Car, error)
}