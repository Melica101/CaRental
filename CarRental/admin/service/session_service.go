package service

import (
	"CarRental/entities"
	"CarRental/admin"
)

//SessionServiceImpl implements admin.SessionService interface
type SessionServiceImpl struct {
	sessionRepo admin.SessionRepository
}

//NewSessionService  returns a new SessionService object
func NewSessionService(sessRepository admin.SessionRepository) admin.SessionService {
	return &SessionServiceImpl{sessionRepo: sessRepository}
}

//Session returns a given stored session
func (ss *SessionServiceImpl) Session(sessionID string) (*entities.Session, []error) {
	return ss.sessionRepo.Session(sessionID)
}

//Sessions Returns all the sessions
func (ss *SessionServiceImpl) Sessions() ([]entities.Session, []error) {
	return ss.sessionRepo.Sessions()
}

//StoreSession stores a given session
func (ss *SessionServiceImpl) StoreSession(session *entities.Session) (*entities.Session, []error) {
	return ss.sessionRepo.StoreSession(session)
}

//DeleteSession deletes a given session
func (ss *SessionServiceImpl) DeleteSession(sessionID string) (*entities.Session, []error) {
	return ss.sessionRepo.DeleteSession(sessionID)
}
