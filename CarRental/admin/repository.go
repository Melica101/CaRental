package admin

import "CarRental/entities"

type AdminRepository interface {
	Admin(id uint) (*entities.Admin, []error)
	StoreAdmin(admin *entities.Admin) (*entities.Admin, []error)
	AdminByEmail(email string) (*entities.Admin, []error)
}

type SessionRepository interface {
	Session(sessionID string) (*entities.Session, []error)
	Sessions() ([]entities.Session, []error)
	StoreSession(session *entities.Session) (*entities.Session, []error)
	DeleteSession(sessionID string) (*entities.Session, []error)
}
