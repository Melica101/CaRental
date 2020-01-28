package repository

import (
	"github.com/jinzhu/gorm"
	"CarRental/entities"
	"CarRental/user"
)

type UserGormRepo struct {
	conn *gorm.DB
}

func NewUserGormRepo(db *gorm.DB) user.UserRepository {
	return &UserGormRepo{conn: db}
}

func (userRepo *UserGormRepo) User(id uint) (*entities.User, []error) {
	user := entities.User{}
	errs := userRepo.conn.First(&user, id).GetErrors()
	return &user, errs

}

func (userRepo *UserGormRepo) StoreUser(user *entities.User) (*entities.User, []error) {
	errs := userRepo.conn.Create(user).GetErrors()
	return user, errs
}
func (userRepo *UserGormRepo) UserByEmail(email string) (*entities.User, []error) {
	user := entities.User{}
	errs := userRepo.conn.First(&user, "email=?", email).GetErrors()
	return &user, errs
}
