package service

import (
	"CarRental/entities"
	"CarRental/user"
)

type UserService struct {
	userRepo user.UserRepository
}

///It's better to return false if error happens instead of allowing user to create a new email when an error occured
func (us *UserService) EmailExists(email string) bool {
	user, error := us.userRepo.UserByEmail(email)
	if user == nil || len(error) > 0 {
		return false
	}
	return true
}

func NewUserService(userRepo user.UserRepository) user.UserService {
	return &UserService{userRepo: userRepo}
}

//get user by id
func (us *UserService) User(id uint) (*entities.User, []error) {
	return us.userRepo.User(id)
}

//get user by Id
func (us *UserService) UserByEmail(email string) (*entities.User, []error) {
	return us.userRepo.UserByEmail(email)
}

//create user
func (us *UserService) StoreUser(user *entities.User) (*entities.User, []error) {
	return us.userRepo.StoreUser(user)
}
