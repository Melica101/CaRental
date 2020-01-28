package services

import (
	"CarRental/entities"
	"CarRental/spam"
)

//SpamService is
type SpamService struct {
	spamRepo spam.SpamRepository
}

//NewSpamService is
func NewSpamService(revoRepo spam.SpamRepository) spam.SpamService {
	return &SpamService{spamRepo: revoRepo}
}

// GetSpams returns all stored Spams
func (rs *SpamService) GetSpams() ([]entities.Spam, []error) {
	revs, errs := rs.spamRepo.GetSpams()
	if len(errs) > 0 {
		return nil, errs
	}
	return revs, errs
}

//GetSpam returns spam by iD
func (rs *SpamService) GetSpam(id uint) (*entities.Spam, []error) {
	rev, errs := rs.spamRepo.GetSpam(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rev, errs
}

// GetSpamsByCarID  retrieves stored comment by its id
func (rs *SpamService) GetSpamsByCarID(CarID uint) ([]entities.Spam, []error) {
	revs, errs := rs.spamRepo.GetSpamsByCarID(CarID)
	if len(errs) > 0 {
		return nil, errs
	}
	return revs, errs
}

// DeleteSpam deletes a given Spam
func (rs *SpamService) DeleteSpam(id uint) (*entities.Spam, []error) {
	revs, errs := rs.spamRepo.DeleteSpam(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return revs, errs
}

//StoreSpam stores spam in db
func (rs *SpamService) StoreSpam(spam *entities.Spam) (*entities.Spam, []error) {
	revs, errs := rs.spamRepo.StoreSpam(spam)
	if len(errs) > 0 {
		return nil, errs
	}
	return revs, errs
}
