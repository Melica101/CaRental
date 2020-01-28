package services

import (
	"CarRental/entities"
	"CarRental/feedback"
)
//FeedbackService is
type FeedbackService struct {
	feedbackRepo feedback.FeedbackRepository
}

func NewFeedbackService(revoRepo feedback.FeedbackRepository) feedback.FeedbackService {
	return &FeedbackService{feedbackRepo: revoRepo}
}

// GetFeedbacks returns all stored feedbacks
func (rs *FeedbackService) GetFeedbacks() ([]entities.Feedback, []error) {
	revs, errs := rs.feedbackRepo.GetFeedbacks()
	if len(errs) > 0 {
		return nil, errs
	}
	return revs, errs
}

func (rs *FeedbackService) GetFeedback(id uint) (*entities.Feedback, []error) {
	rev, errs := rs.feedbackRepo.GetFeedback(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rev, errs
}

func (rs *FeedbackService) StoreFeedback(feedback *entities.Feedback) (*entities.Feedback, []error) {
	revs, errs := rs.feedbackRepo.StoreFeedback(feedback)
	if len(errs) > 0 {
		return nil, errs
	}
	return revs, errs
}
