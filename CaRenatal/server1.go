package main

import (
	"database/sql"
	"net/http"

	"github.com/amalica/CarRental/car"
	"github.com/amalica/CarRental/user"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	db, err := sql.Open("mysql", "root:0911@tcp(localhost:3306)/carrental?parseTime=true")
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	user.RepositoryGlobal = user.NewRepository(db)
	user.ServiceGlobal = user.NewService(user.RepositoryGlobal)

	car.RepositoryGlobal = car.NewRepository(db)
	car.ServiceGlobal = car.NewService(car.RepositoryGlobal)

	mux := http.NewServeMux()

	// fileServer1 := http.FileServer(http.Dir("assets"))
	// mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer1))

	mux.HandleFunc("/Register", user.Register)
	mux.HandleFunc("/Login", user.Login)
	mux.HandleFunc("/Logout", user.Logout)
	mux.HandleFunc("/PostCar", car.PostCar)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
