package repository

import (
	"github.com/jinzhu/gorm"
	"CarRental/entities"
	"CarRental/spam"
)

//SpamGormRepo is
type SpamGormRepo struct {
	conn *gorm.DB
}

//NewSpamGormRepo returns new object of SpamGormRepo
func NewSpamGormRepo(db *gorm.DB) spam.SpamRepository {
	return &SpamGormRepo{conn: db}
}

//GetSpams returns all Spams stored in the database
func (revRepo *SpamGormRepo) GetSpams() ([]entities.Spam, []error) {
	revws := []entities.Spam{}
	errs := revRepo.conn.Find(&revws).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return revws, errs
}

//GetSpamsByCarID retrieve a Spam from the database by its Carid
func (revRepo *SpamGormRepo) GetSpamsByCarID(CarID uint) ([]entities.Spam, []error) {
	revo := []entities.Spam{}
	errs := revRepo.conn.Set("gorm:auto_preload", true).Find(&revo, "Car_id=?", CarID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return revo, errs
}

//GetSpam returns spam by ID
func (revRepo *SpamGormRepo) GetSpam(id uint) (*entities.Spam, []error) {
	spam := entities.Spam{}
	errs := revRepo.conn.Set("gorm:auto_preload", true).First(&spam, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &spam, errs
}

// DeleteSpam deletes a given Spam from the database
func (revRepo *SpamGormRepo) DeleteSpam(id uint) (*entities.Spam, []error) {
	spam, errs := revRepo.GetSpam(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = revRepo.conn.Delete(spam, spam.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return spam, errs
}

// StoreSpam stores a given review in the database
func (revRepo *SpamGormRepo) StoreSpam(spam *entities.Spam) (*entities.Spam, []error) {
	revi := spam
	errs := revRepo.conn.Create(revi).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return revi, errs
}
