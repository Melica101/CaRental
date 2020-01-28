package repository

import (
	"CarRental/entities"
	"CarRental/feedback"

	"github.com/jinzhu/gorm"
)

//FeedbackGormRepo is
type FeedbackGormRepo struct {
	conn *gorm.DB
}

//NewFeedbackGormRepo returns new object of feedbackGormRepo
func NewFeedbackGormRepo(db *gorm.DB) feedback.FeedbackRepository {
	return &FeedbackGormRepo{conn: db}
}

//GetFeedbacks returns all feedbacks stored in the database
func (revRepo *FeedbackGormRepo) GetFeedbacks() ([]entities.Feedback, []error) {
	revws := []entities.Feedback{}
	errs := revRepo.conn.Find(&revws).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return revws, errs
}

//GetFeedback returns feedbacks by ID
func (revRepo *FeedbackGormRepo) GetFeedback(id uint) (*entities.Feedback, []error) {
	feedback := entities.Feedback{}
	errs := revRepo.conn.Set("gorm:auto_preload", true).First(&feedback, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &feedback, errs
}

// StoreFeedback stores a given review in the database
func (revRepo *FeedbackGormRepo) StoreFeedback(feedback *entities.Feedback) (*entities.Feedback, []error) {
	revi := feedback
	errs := revRepo.conn.Create(revi).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return revi, errs
}
