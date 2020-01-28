package user

import "CarRental/entities"

type UserRepository interface {
	User(id uint) (*entities.User, []error)
	StoreUser(user *entities.User) (*entities.User, []error)
	UserByEmail(email string) (*entities.User, []error)
}
type RoleRepository interface {
	Roles() ([]entities.Role, []error)
	Role(id uint) (*entities.Role, []error)
	RoleByName(name string) (*entities.Role, []error)
	UpdateRole(role *entities.Role) (*entities.Role, []error)
	DeleteRole(id uint) (*entities.Role, []error)
	StoreRole(role *entities.Role) (*entities.Role, []error)
}
type SessionRepository interface {
	Session(sessionID string) (*entities.Session, []error)
	Sessions() ([]entities.Session, []error)
	StoreSession(session *entities.Session) (*entities.Session, []error)
	DeleteSession(sessionId string) (*entities.Session, []error)
}
