package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Car is car structure
type Car struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null"`
	CarModel     string
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

//Admin is admin struct
type Admin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Password string
}

//Role is role struct
type Role struct {
	ID uint
	Name string `gorm:"type:varchar(255)"`
}
 
//Session is session struct
type Session struct {
	gorm.Model
	SessionID  string `gorm:"type:varchar(255);not null"`
	UUID       uint
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}

//User is user struct
type User struct {
	gorm.Model
	Firstname   string `gorm:"type:varchar(255);not null"`
	Lastname    string `gorm:"type:varchar(255);not null"`
	Password    string `gorm:"type:varchar(255)"`
	Phonenumber string `gorm:"type:varchar(16);not null; unique"`
	Email       string `gorm:"type:varchar(255);not null;unique"`
	Address     string
	City        string
	Accountnum  string
	RoleID uint
}

//Feedback is feedback struct
type Feedback struct {
	gorm.Model
	UserID   uint
	Feedback string
}

//Spam is spam struct
type Spam struct {
	gorm.Model
	UserID uint
	CarID  uint
}
