package repository

import (
	"CarRental/car"
	"CarRental/entities"

	"github.com/jinzhu/gorm"
)

//CarGormRepo
type CarGormRepo struct {
	conn *gorm.DB
}

//NewCarGormRepo -
func NewCarGormRepo(db *gorm.DB) car.CarRepository {
	return &CarGormRepo{conn: db}
}

//GetCars returns all cars stored in the database
func (carRepo *CarGormRepo) GetCars() ([]entities.Car, []error) {
	shps := []entities.Car{}
	errs := carRepo.conn.Find(&shps).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return shps, errs
}

//GetCar -
func (carRepo *CarGormRepo) GetCar(id uint) (*entities.Car, []error) {
	car := entities.Car{}
	errs := carRepo.conn.First(&car, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &car, errs
}

// UpdateCar updates a given car in the database
func (carRepo *CarGormRepo) UpdateCar(car *entities.Car) (*entities.Car, []error) {
	shps := car
	errs := carRepo.conn.Save(shps).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return shps, errs
}

// DeleteCar deletes a given car from the database
func (carRepo *CarGormRepo) DeleteCar(id uint) (*entities.Car, []error) {
	shps, errs := carRepo.GetCar(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = carRepo.conn.Delete(shps, shps.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return shps, errs
}

// StoreCar stores a given car in the database
func (carRepo *CarGormRepo) StoreCar(car *entities.Car) (*entities.Car, []error) {
	shps := car
	errs := carRepo.conn.Create(shps).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return shps, errs
}

//GetCarByUserID gets car by user ID
func (carRepo *CarGormRepo) GetCarByUserID(userID uint) (*entities.Car, []error) {
	car := entities.Car{}
	errs := carRepo.conn.Set("gorm:auto_preload", true).Find(&car, "user_id=?", userID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &car, errs
}
