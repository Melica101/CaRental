package main

import (
	"CarRental/entities"
	"CarRental/rtoken"
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	//carRepoImport "CarRental/car/repository"
	//carServiceImport "CarRental/car/service"
	"CarRental/carental/http/handler"
	spamRepoImport "CarRental/spam/repository"
	spamServiceImport "CarRental/spam/services"
	userRepoImport "CarRental/user/repository"
	userServiceImport "CarRental/user/service"
)

func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.DropTable(
		&entities.User{},
		&entities.Spam{},
		&entities.Feedback{},
		&entities.Role{},
		&entities.Admin{},
		&entities.Session{},
		&entities.Car{},
	).GetErrors()

	errs = dbconn.CreateTable(
		&entities.User{},
		&entities.Spam{},
		&entities.Feedback{},
		&entities.Role{},
		&entities.Admin{},
		&entities.Session{},
		&entities.Car{},
	).GetErrors()
	errs = dbconn.Create(&entities.Role{ID: 1, Name: "USER"}).GetErrors()
	errs = dbconn.Create(&entities.Role{ID: 2, Name: "ADMIN"}).GetErrors()
	if errs != nil {
		return errs
	}

	return nil
}

func main() {
	tmpl := template.Must(template.ParseGlob("../ui/templates/*"))
	dbconn, err := gorm.Open("postgres", "postgres://postgres:root@localhost/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	//createTables(dbconn)

	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	userRepo := userRepoImport.NewUserGormRepo(dbconn)
	userService := userServiceImport.NewUserService(userRepo)

	sessionRepo := userRepoImport.NewSessionGormRepo(dbconn)
	sessionService := userServiceImport.NewSessionService(sessionRepo)

	roleRepo := userRepoImport.NewRoleGormRepo(dbconn)
	roleService := userServiceImport.NewRoleService(roleRepo)

	//carRepo := carRepoImport.NewCarGormRepo(dbconn)
	//carService := carServiceImport.NewCarService(carRepo)

	spamRepo := spamRepoImport.NewSpamGormRepo(dbconn)
	spamService := spamServiceImport.NewSpamService(spamRepo)

	userHandler := handler.NewUserHandler(tmpl, userService, sessionService, roleService, csrfSignKey)
	adminDashboardHandler := handler.NewAdminDashboardHandler(tmpl, spamService, csrfSignKey)

	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", userHandler.Index)
	http.Handle("/admin", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminDashboardHandler.AdminIndex))))
	http.HandleFunc("/login", userHandler.Login)
	//http.Handle("/signup", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminDashboardHandler.AdminSignUp))))
	//http.Handle("/basicInfo", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminDashboardHandler.AdminBasicInfo))))
	http.Handle("/logout", userHandler.Authenticated(http.HandlerFunc(userHandler.Logout)))
	http.HandleFunc("/signup", userHandler.SignUp)
	http.ListenAndServe(":8181", nil)
}
