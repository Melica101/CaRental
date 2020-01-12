package user

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/amalica/CarRental/entities"
)

// Repository is a struct that define the Repository type.
type Repository struct {
	connection *sql.DB
}

// IRepository is an interface that specifies database operations on User type.
type IRepository interface {
	AddUser(*entities.User) error
	GetUser(key string) *entities.User
	CreateSession(email string) *Session
	RemoveSession(sessionID string)
	GetEmail(sessionID string) string
}

// NewRepository is a function that return new IRepository type.
func NewRepository(conn *sql.DB) IRepository {
	return &Repository{connection: conn}
}

// AddUser is a method that adds a user to the provided database.
func (psql *Repository) AddUser(user *entities.User) error {

	stmt, err := psql.connection.Prepare(`INSERT INTO users (Firstname, Lastname, Password, Phonenumber, Email,
		Address, City) VALUES (?,?,?,?,?,?,?)`)

	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(
		user.Firstname,
		user.Lastname,
		user.Password,
		user.Phonenumber,
		user.Email,
		user.Address,
		user.City)

	if err != nil {
		return err
	}
	return nil
}

// GetUser is a method that search a user a key and returns the user
func (psql *Repository) GetUser(key string) *entities.User {

	stmt, err := psql.connection.Prepare("SELECT * FROM users WHERE Email=? || ID=?")
	row := stmt.QueryRow(key, key)
	if err != nil {
		panic(err)
	}

	var user entities.User
	user.ID = -1

	row.Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Password,
		&user.Phonenumber,
		&user.Email,
		&user.Address,
		&user.City)

	return &user
}

// CreateSession is a method that returns a newly created session.
func (psql *Repository) CreateSession(email string) *Session {

	stmt, err := psql.connection.Prepare(`INSERT INTO sessions (ID, Email, LoggedIn) VALUES (?,?,?)`)

	if err != nil {
		panic(err)
	}

	session := Session{
		SessionID: RandomStringGN(),
		Email:     email,
		LoggedIn:  true,
	}
	_, err = stmt.Exec(session.SessionID, session.Email, session.LoggedIn)
	if err != nil {
		return nil
	}

	return &session
}

// RemoveSession is a method that deletes a user session from the database.
func (psql *Repository) RemoveSession(sessionID string) {

	stmt, err := psql.connection.Prepare(`DELETE FROM sessions WHERE ID=?`)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(sessionID)

	if err != nil {
		panic(err)
	}
}

// GetEmail is a method that returns an email from the session.
func (psql *Repository) GetEmail(sessionID string) string {

	var email string
	stmt, err := psql.connection.Prepare("SELECT Email FROM sessions WHERE ID=?")
	row := stmt.QueryRow(sessionID)
	if err != nil {
		panic(err)
	}
	row.Scan(&email)

	return email

}

// RandomStringGN is a function that generates random string every time.
func RandomStringGN() string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
