package service

import (
	"CarRental/entities"
	"CarRental/user"
)

// RoleService implements menu.RoleService interface
type RoleService struct {
	roleRepo user.RoleRepository
}

func NewRoleService(RoleRepo user.RoleRepository) user.RoleService {
	return &RoleService{roleRepo: RoleRepo}
}

func (rs *RoleService) Roles() ([]entities.Role, []error) {
	return rs.roleRepo.Roles()
}

func (rs *RoleService) RoleByName(name string) (*entities.Role, []error) {
	return rs.roleRepo.RoleByName(name)
}

func (rs *RoleService) Role(id uint) (*entities.Role, []error) {
	return rs.roleRepo.Role(id)
}

func (rs *RoleService) UpdateRole(role *entities.Role) (*entities.Role, []error) {
	return rs.roleRepo.UpdateRole(role)
}

func (rs *RoleService) DeleteRole(id uint) (*entities.Role, []error) {
	return rs.roleRepo.DeleteRole(id)
}

func (rs *RoleService) StoreRole(role *entities.Role) (*entities.Role, []error) {
	return rs.roleRepo.StoreRole(role)
}
