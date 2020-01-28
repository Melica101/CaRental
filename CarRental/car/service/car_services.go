package service

import (
	"CarRental/car"
	"CarRental/entities"
)

//CarService -
type CarService struct {
	carRepo car.CarRepository
}

//GetCarByUserID -
func (sp *CarService) GetCarByUserID(userID uint) (*entities.Car, []error) {
	shp, errs := sp.carRepo.GetCarByUserID(userID)
	if len(errs) > 0 {
		return nil, errs
	}
	return shp, errs
}

//NewCarService -
func NewCarService(shpRepo car.CarRepository) car.CarService {
	return &CarService{carRepo: shpRepo}
}

// GetCars returns all stored cars
func (sp *CarService) GetCars() ([]entities.Car, []error) {
	shp, errs := sp.carRepo.GetCars()
	if len(errs) > 0 {
		return nil, errs
	}
	return shp, errs
}

//GetCar gets car by Id
func (sp *CarService) GetCar(id uint) (*entities.Car, []error) {
	car, errs := sp.carRepo.GetCar(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return car, errs
}

// UpdateCar updates a given car in a database
func (sp *CarService) UpdateCar(car *entities.Car) (*entities.Car, []error) {
	shp, errs := sp.carRepo.UpdateCar(car)
	if len(errs) > 0 {
		return nil, errs
	}
	return shp, errs
}

// DeleteCar deletes a given car
func (sp *CarService) DeleteCar(id uint) (*entities.Car, []error) {
	shp, errs := sp.carRepo.DeleteCar(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return shp, errs
}

// StoreCar stores a given car
func (sp *CarService) StoreCar(car *entities.Car) (*entities.Car, []error) {
	shp, errs := sp.carRepo.StoreCar(car)
	if len(errs) > 0 {
		return nil, errs
	}
	return shp, errs
}
