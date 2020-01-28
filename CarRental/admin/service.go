package admin

import "CarRental/entities"

type AdminService interface {
	Admin(id uint) (*entities.Admin, []error)
	StoreAdmin(admin *entities.Admin) (*entities.Admin, []error)
	AdminByEmail(email string) (*entities.Admin, []error)
	EmailExists(email string) bool
}

// SessionService specifies logged in admin session related service
type SessionService interface {
	Session(sessionID string) (*entities.Session, []error)
	Sessions() ([]entities.Session, []error)
	StoreSession(session *entities.Session) (*entities.Session, []error)
	DeleteSession(sessionID string) (*entities.Session, []error)
}
