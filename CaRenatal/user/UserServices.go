package user

import (
	"errors"

	"github.com/amalica/CarRental/entities"
)

// Service is a struct that defines the UserService type.
type Service struct {
	conn IRepository
}

// IService is an interface that specifies what a User type can do.
type IService interface {
	RegisterUser(*entities.User) error
	Login(email string, password string) (*entities.User, error)
	SearchUser(key string) (*entities.User, error)
}

// NewService is a function that returns a new UserService type.
func NewService(connection IRepository) IService {
	return &Service{conn: connection}
}

// RegisterUser is a method that register a new user to the system.
func (service *Service) RegisterUser(user *entities.User) error {
	err := service.conn.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}

// Login is a method that returns a user if enter a valid user name and password.
func (service *Service) Login(email string, password string) (*entities.User, error) {
	user := service.conn.GetUser(email)
	if user.ID == -1 {
		return nil, errors.New("invalid user email")
	}

	if user.Password != password {
		return nil, errors.New("invalid user password")
	}

	return user, nil

}

// SearchUser is a method that searchs a user using a provided key and return user if found.
func (service *Service) SearchUser(key string) (*entities.User, error) {

	user := service.conn.GetUser(key)
	if user.ID == -1 {
		return user, errors.New("user not found")
	}

	return user, nil

}
