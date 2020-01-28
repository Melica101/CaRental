package service

import (
	"CarRental/entities"
	"CarRental/admin"
)

//AdminService -
type AdminService struct {
	adminRepo admin.AdminRepository
}

//EmailExists It's better to return false if error happens instead of allowing admin to create a new email when an error occured
func (us *AdminService) EmailExists(email string) bool {
	admin, error := us.adminRepo.AdminByEmail(email)
	if admin == nil || len(error) > 0 {
		return false
	}
	return true
}

//NewAdminService -
func NewAdminService(adminRepo admin.AdminRepository) admin.AdminService {
	return &AdminService{adminRepo: adminRepo}
}

//Admin get admin by id
func (us *AdminService) Admin(id uint) (*entities.Admin, []error) {
	return us.adminRepo.Admin(id)
}

//AdminByEmail get admin by email
func (us *AdminService) AdminByEmail(email string) (*entities.Admin, []error) {
	return us.adminRepo.AdminByEmail(email)
}

//StoreAdmin create admin
func (us *AdminService) StoreAdmin(admin *entities.Admin) (*entities.Admin, []error) {
	return us.adminRepo.StoreAdmin(admin)
}
