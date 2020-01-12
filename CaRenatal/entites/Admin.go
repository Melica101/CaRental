package entities

// Admin is a struct that defines the type Admin.
type Admin struct {
	ID       int
	Name     string
	Password string
}

// NewAdmin is a function that return a new Admin type from provided arguments.
func NewAdmin(name, password string) *Admin {
	Admin := Admin{
		Name:     name,
		Password: password,
	}
	return &Admin
}
