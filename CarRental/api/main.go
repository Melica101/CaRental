package main

import (
	handler "CarRental/carental/http/handler"
	"CarRental/entities"
	"net/http"

	carRepoImport "CarRental/car/repository"
	carServiceImport "CarRental/car/service"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.DropTable(
		&entities.Car{},
		&entities.Spam{},
		&entities.Feedback{},
		&entities.Role{},
		&entities.Admin{},
		&entities.Session{},
		&entities.Car{},
	).GetErrors()

	errs = dbconn.CreateTable(
		&entities.Car{},
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
	dbconn, err := gorm.Open("postgres", "postgres://postgres:root@localhost/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	//createTables(dbconn)

	carRepo2 := carRepoImport.NewCarGormRepo(dbconn)
	carService2 := carServiceImport.NewCarService(carRepo2)
	APIRouter := httprouter.New()
	APIRouter.ServeFiles("/assets/*filepath", http.Dir("../ui/assets"))
	CarHandler := handler.NewCarHandler(carService2)
	APIRouter.GET("/api/cars", CarHandler.GetCars)
	APIRouter.GET("/api/cars/:id", CarHandler.GetSinglecar)
	APIRouter.DELETE("/api/cars/:id", CarHandler.DeleteCar)
	APIRouter.PUT("/api/cars/:id", CarHandler.UpdateCar)
	APIRouter.POST("/api/cars", CarHandler.PostCar)
	http.ListenAndServe(":8181", APIRouter)
}
