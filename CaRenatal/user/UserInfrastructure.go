package user

import (
	"errors"
	"net/http"

	"github.com/amalica/CarRental/entities"
)

var (
	// RepositoryGlobal is a global variable
	RepositoryGlobal IRepository
	// ServiceGlobal is a global variable
	ServiceGlobal IService
)

// Session is a struct that holds the session inforamtion.
type Session struct {
	SessionID string
	Email     string
	LoggedIn  bool
}

// Register is a Handler function that takes argument from the form and pass it to service.
func Register(w http.ResponseWriter, r *http.Request) {

	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	phonenumber := r.FormValue("phonenumber")
	email := r.FormValue("email")
	address := r.FormValue("address")
	city := r.FormValue("city")
	password := r.FormValue("password")

	user := entities.NewUser(firstname, lastname, email, password, phonenumber, address, city)
	err := ServiceGlobal.RegisterUser(user)
	if err != nil {
		panic(err)
	}

}

// Login is a Handler fucntion that
func Login(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := ServiceGlobal.Login(email, password)
	if err != nil {
		panic(err.Error())
	}

	session := RepositoryGlobal.CreateSession(user.Email)
	if err != nil {
		panic(err)
	}

	cookie := http.Cookie{
		Name:     "user",
		Value:    session.SessionID,
		MaxAge:   3600 * 3,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

}

// Logout is a Handler function that logs out a user.
func Logout(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("user")
	if err != nil {
		panic(err)
	}
	sessionID := cookie.Value
	RepositoryGlobal.RemoveSession(sessionID)
}

// AuthenticateUser is a Handler function that checks if a user is logged or not.
func AuthenticateUser(w http.ResponseWriter, r *http.Request) (*entities.User, error) {
	cookie, err := r.Cookie("user")
	if err != nil {
		panic(err)
	}
	sessionID := cookie.Value
	email := RepositoryGlobal.GetEmail(sessionID)
	user, err := ServiceGlobal.SearchUser(email)
	if err != nil {
		return nil, errors.New("user not logged in")
	}
	return user, nil

}
