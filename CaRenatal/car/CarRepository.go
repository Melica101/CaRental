package car

import (
	"database/sql"
	"fmt"

	"github.com/amalica/CarRental/entities"
)

// Repository is a struct that define the Repository type.
type Repository struct {
	connection *sql.DB
}

// IRepository is an interface that specifies database operations on Car type.
type IRepository interface {
	AddCar(*entities.Car) (string, error)
	GetCar(string) *entities.Car
	CountCars() (totalNumOfMembers int)
	AddPostedCar(carID string, userID int) error
}

// NewRepository is a function that return new Repository type.
func NewRepository(conn *sql.DB) IRepository {
	return &Repository{connection: conn}
}

// AddCar is a method that adds a Car to the provided database.
func (psql *Repository) AddCar(Car *entities.Car) (string, error) {

	totalNumOfCars := psql.CountCars()
	Car.ID = fmt.Sprintf("CarID%d", totalNumOfCars+1)

	stmt, err := psql.connection.Prepare(`INSERT INTO cars (
		Color,
		FuelType,
		FuelUsage,
		Model,
		ID,
		Name,
		Photo,
		PlateNumber,
		Price,
		SeatsNumber,
		Transmission,
		Year
	  )
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(
		Car.Color,
		Car.FuelType,
		Car.FuelUsage,
		Car.Model,
		Car.ID,
		Car.Name,
		Car.Photo,
		Car.PlateNumber,
		Car.Price,
		Car.SeatsNumber,
		Car.Transmission,
		Car.Year)

	if err != nil {
		return "", err
	}
	return Car.ID, nil
}

// GetCar is a method that searches and deliver the needed Car using the Car id.
func (psql *Repository) GetCar(pid string) *entities.Car {

	var Car entities.Car
	stmt, err := psql.connection.Prepare("SELECT * FROM cars WHERE ID=?")
	if err != nil {
		return nil
	}
	row := stmt.QueryRow(pid)

	row.Scan(
		&Car.ID,
		&Car.Name,
		&Car.Model,
		&Car.Year,
		&Car.Color,
		&Car.Transmission,
		&Car.FuelType,
		&Car.FuelUsage,
		&Car.SeatsNumber,
		&Car.PlateNumber,
		&Car.Price,
		&Car.Photo,
	)

	return &Car

}

// AddPostedCar is a method that adds a posted car to posted_cars table and link it to the owner(user).
func (psql *Repository) AddPostedCar(carID string, userID int) error {

	stmt, err := psql.connection.Prepare(`INSERT INTO posted_cars (carID, userID) VALUES (?,?)`)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(carID, userID)

	if err != nil {
		return err
	}
	return nil

}

// CountCars is a method that is used for counting the member of a table cars.
func (psql *Repository) CountCars() (totalNumOfMembers int) {

	stmt, err := psql.connection.Prepare("SELECT COUNT(*) FROM cars")
	if err != nil {
		return
	}
	row := stmt.QueryRow()
	row.Scan(&totalNumOfMembers)
	return

}
