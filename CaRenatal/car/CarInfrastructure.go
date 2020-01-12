package car

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/amalica/CarRental/entities"
	"github.com/amalica/CarRental/user"
)

var (
	// RepositoryGlobal is a global variable
	RepositoryGlobal IRepository
	// ServiceGlobal is a global variable
	ServiceGlobal IService
)

//PostCar is a method that enables user to post cars.
func PostCar(w http.ResponseWriter, r *http.Request) {

	LUser, err := user.AuthenticateUser(w, r)
	if err != nil {
		panic(err)
	}

	name := r.FormValue("name")
	model := r.FormValue("model")
	yearString := r.FormValue("year")
	color := r.FormValue("color")
	transmission := r.FormValue("transmission")
	fuelType := r.FormValue("fuelType")
	fuelUsage := r.FormValue("fuelUsage")
	seatsNumberString := r.FormValue("seatsNumber")
	plateNumber := r.FormValue("plateNumber")
	priceString := r.FormValue("price")
	photo := r.FormValue("photo")

	price, _ := strconv.ParseFloat(priceString, 0)
	seatsNumber, _ := strconv.ParseInt(seatsNumberString, 0, 0)

	year, _ := time.Parse("01/02/2006", yearString)

	fm, fh, err := r.FormFile("photo")
	defer fm.Close()

	path, _ := os.Getwd()
	suffix := ""
	endPoint := 0

	for i := len(fh.Filename) - 1; i >= 0; i-- {
		if fh.Filename[i] == '.' {
			endPoint = i
			break
		}
	}

	for ; endPoint < len(fh.Filename); endPoint++ {
		suffix += string(fh.Filename[endPoint])
	}

	NewFileName := fmt.Sprintf("carPic_"+name+"_%s"+suffix, user.RandomStringGN())
	path = filepath.Join(path, "assets", NewFileName)

	out, _ := os.Create(path)
	defer out.Close()

	_, err = io.Copy(out, fm)

	if err != nil {
		panic(err)
	}

	photo = NewFileName

	car := entities.NewCar(name, model, color, transmission, fuelType, fuelUsage, plateNumber, photo, seatsNumber, price, year)
	if err := ServiceGlobal.PostCar(car, LUser.ID); err != nil {
		panic(err)
	}

}
