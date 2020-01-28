package util

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func ParseForm(w http.ResponseWriter,r *http.Request) bool{
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return false
	}
	return true
}

func HashPassword(password string) ([]byte,error){
	return bcrypt.GenerateFromPassword([]byte(password), 12)
}
func ArePasswordsSame(hashedPassword string,rawPassword string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	return err != bcrypt.ErrMismatchedHashAndPassword
}