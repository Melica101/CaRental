package spam

import (
	"CarRental/entities"
)

type SpamRepository interface {
	GetSpams() ([]entities.Spam, []error)
	GetSpam(id uint) (*entities.Spam, []error)
	DeleteSpam(id uint) (*entities.Spam, []error)
	GetSpamsByCarID(CarID uint) ([]entities.Spam, []error)
	StoreSpam(spam *entities.Spam) (*entities.Spam, []error)
}
