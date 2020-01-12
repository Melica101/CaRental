package entities

// User is a struct that defines the type User.
type User struct {
	ID          int
	Firstname   string
	Lastname    string
	Password    string
	Phonenumber string
	Email       string
	Address     string
	City        string
}

// NewUser is a function that return a new User type from provided arguments.
func NewUser(firstname, lastname, email, password, phonenumber, Address, city string) *User {
	user := User{
		Firstname:   firstname,
		Lastname:    lastname,
		Password:    password,
		Phonenumber: phonenumber,
		Email:       email,
		Address:     Address,
		City:        city,
	}
	return &user
}
