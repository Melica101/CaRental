package repository

import (
	"github.com/jinzhu/gorm"
	"CarRental/entities"
	"CarRental/admin"
)

// SessionGormRepo implements admin.SessionRepository interface
type SessionGormRepo struct {
	conn *gorm.DB
}

// NewSessionGormRepo  returns a new SessionGormRepo object
func NewSessionGormRepo(db *gorm.DB) admin.SessionRepository {
	return &SessionGormRepo{conn: db}
}

// Session returns a given stored session
func (sr *SessionGormRepo) Session(sessionId string) (*entities.Session, []error) {
	session := entities.Session{}
	errs := sr.conn.Find(&session, "session_id=?", sessionId).GetErrors()
	return &session, errs
}

// Returns all the sessions
func (sr *SessionGormRepo) Sessions() ([]entities.Session, []error) {
	sessions := []entities.Session{}
	errs := sr.conn.Find(&sessions).GetErrors()
	return sessions, errs
}

// StoreSession stores a given session
func (sr *SessionGormRepo) StoreSession(session *entities.Session) (*entities.Session, []error) {
	sess := session
	errs := sr.conn.Save(sess).GetErrors()
	return sess, errs
}

// DeleteSession deletes a given session
func (sr *SessionGormRepo) DeleteSession(sessionId string) (*entities.Session, []error) {
	sess, errs := sr.Session(sessionId)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = sr.conn.Delete(sess, sessionId).GetErrors()
	return sess, errs
}
