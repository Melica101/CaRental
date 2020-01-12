package car

import "github.com/amalica/CarRental/entities"

// Service is a struct that defines the Car Service type.
type Service struct {
	conn IRepository
}

// IService is an interface that specifies what a Car type can do.
type IService interface {
	PostCar(car *entities.Car, userID int) error
}

// NewService is a function that returns a new UserService type.
func NewService(connection IRepository) IService {
	return &Service{conn: connection}
}

// PostCar is a method that enables car posting service.
func (service *Service) PostCar(car *entities.Car, userID int) error {

	carID, err := service.conn.AddCar(car)
	if err != nil {
		return err
	}

	err = service.conn.AddPostedCar(carID, userID)
	if err != nil {
		return err
	}

	return nil
}
