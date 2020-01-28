package car

import "CarRental/entities"

type CarRepository interface {
	GetCars() ([]entities.Car, []error)
	GetCar(id uint) (*entities.Car, []error)
	StoreCar(car *entities.Car) (*entities.Car, []error)
	UpdateCar(car *entities.Car) (*entities.Car, []error)
	DeleteCar(id uint) (*entities.Car, []error)
	GetCarByUserID(userID uint) (*entities.Car, []error)
}
