package user

import "CarRental/entities"

type UserService interface {
	User(id uint) (*entities.User, []error)
	StoreUser(user *entities.User) (*entities.User, []error)
	UserByEmail(email string) (*entities.User, []error)
	EmailExists(email string) bool
}
type RoleService interface {
	Roles() ([]entities.Role, []error)
	Role(id uint) (*entities.Role, []error)
	RoleByName(name string) (*entities.Role, []error)
	UpdateRole(role *entities.Role) (*entities.Role, []error)
	DeleteRole(id uint) (*entities.Role, []error)
	StoreRole(role *entities.Role) (*entities.Role, []error)
}

// SessionService specifies logged in user session related service
type SessionService interface {
	Session(sessionId string) (*entities.Session, []error)
	Sessions() ([]entities.Session, []error)
	StoreSession(session *entities.Session) (*entities.Session, []error)
	DeleteSession(sessionId string) (*entities.Session, []error)
}
