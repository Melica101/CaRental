package repository

import (
	"github.com/jinzhu/gorm"
	"CarRental/entities"
	"CarRental/search"
)

//SearchGormRepo -
type SearchGormRepo struct {
	conn *gorm.DB
}

//GetByName -
func (s SearchGormRepo) GetByName(keyword string) ([]entities.Car, error) {
	cars := []entities.Car{}
	s.conn.Where("name like ?", "%"+keyword+"%").Find(&cars)
	return cars, nil
}

//NewSearchGormRepo -
func NewSearchGormRepo(db *gorm.DB) search.SearchRepository {
	return &SearchGormRepo{conn: db}
}