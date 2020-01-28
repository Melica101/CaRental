package feedback

import (
	"CarRental/entities"
)

type FeedbackService interface {
	GetFeedbacks() ([]entities.Feedback, []error)
	GetFeedback(id uint) (*entities.Feedback, []error)
	StoreFeedback(feedback *entities.Feedback) (*entities.Feedback, []error)
}
