package service

import (
	"CarRental/entities"
	"CarRental/search"
)

//SearchService -
type SearchService struct {
	searchRepo search.SearchRepository
}

//GetByName -
func (s SearchService) GetByName(keyword string) ([]entities.Car, error) {
	return s.searchRepo.GetByName(keyword)
}

//NewSearchService -
func NewSearchService(searchRepo search.SearchRepository) search.SearchService {
	return &SearchService{searchRepo: searchRepo}
}