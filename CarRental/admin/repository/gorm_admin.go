package repository

import (
	"github.com/jinzhu/gorm"
	"CarRental/entities"
	"CarRental/admin"
)

type AdminGormRepo struct {
	conn *gorm.DB
}

func NewAdminGormRepo(db *gorm.DB) admin.AdminRepository {
	return &AdminGormRepo{conn: db}
}

func (adminRepo *AdminGormRepo) Admin(id uint) (*entities.Admin, []error) {
	admin := entities.Admin{}
	errs := adminRepo.conn.First(&admin, id).GetErrors()
	return &admin, errs

}

func (adminRepo *AdminGormRepo) StoreAdmin(admin *entities.Admin) (*entities.Admin, []error) {
	errs := adminRepo.conn.Create(admin).GetErrors()
	return admin, errs
}
func (adminRepo *AdminGormRepo) AdminByEmail(email string) (*entities.Admin, []error) {
	admin := entities.Admin{}
	errs := adminRepo.conn.First(&admin, "email=?", email).GetErrors()
	return &admin, errs
}
