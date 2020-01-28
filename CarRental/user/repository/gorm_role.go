package repository

import (
	"github.com/jinzhu/gorm"
	"CarRental/entities"
	"CarRental/user"
)

type RoleGormRepo struct {
	conn *gorm.DB
}

func NewRoleGormRepo(db *gorm.DB) user.RoleRepository {
	return &RoleGormRepo{conn: db}
}

func (roleRepo *RoleGormRepo) Role(id uint) (*entities.Role, []error) {
	role := entities.Role{}
	errs := roleRepo.conn.First(&role, id).GetErrors()
	return &role, errs
}

func (roleRepo *RoleGormRepo) Roles() ([]entities.Role, []error) {
	roles := []entities.Role{}
	errs := roleRepo.conn.Find(&roles).GetErrors()
	return roles, errs
}

func (roleRepo *RoleGormRepo) RoleByName(name string) (*entities.Role, []error) {
	role := entities.Role{}
	errs := roleRepo.conn.Find(&role, "name=?", name).GetErrors()
	return &role, errs
}

func (roleRepo *RoleGormRepo) StoreRole(role *entities.Role) (*entities.Role, []error) {
	errs := roleRepo.conn.Create(&role).GetErrors()
	return role, errs
}

func (roleRepo *RoleGormRepo) UpdateRole(role *entities.Role) (*entities.Role, []error) {
	r := role
	errs := roleRepo.conn.Save(r).GetErrors()
	return r, errs
}

func (roleRepo *RoleGormRepo) DeleteRole(id uint) (*entities.Role, []error) {
	r, errs := roleRepo.Role(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = roleRepo.conn.Delete(r, id).GetErrors()
	return r, errs
}
