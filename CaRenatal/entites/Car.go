package entities

import "time"

// Car is a car structure.
type Car struct {
	ID           string
	Name         string
	Model        string
	Year         time.Time
	Color        string
	Transmission string
	FuelType     string
	FuelUsage    string
	SeatsNumber  int64
	PlateNumber  string
	Price        float64
	Photo        string
}

// NewCar is a function that returns a new Car type from provided arguments.
func NewCar(name, model, color, transmission, fuelType, fuelUsage, plateNumber, photo string, seatsNumber int64, price float64, year time.Time) *Car {
	car := Car{
		Name:         name,
		Model:        model,
		Year:         year,
		Color:        color,
		Transmission: transmission,
		FuelType:     fuelType,
		FuelUsage:    fuelUsage,
		SeatsNumber:  seatsNumber,
		PlateNumber:  plateNumber,
		Price:        price,
		Photo:        photo}

	return &car

}
