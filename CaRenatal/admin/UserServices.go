package user

import (
	"errors"
	"math/rand"
	"net/smtp"
	"regexp"
	"time"

	"github.com/amalica/CarRental/entities"
)

// Service is a struct that defines the UserService type.
type Service struct {
	conn IRepository
}

// Identification is a struct that hold a basic information about third party authentication.
type Identification struct {
	TpFlag          bool
	From            string
	ConfirmPassword string
}

// TempVerificationPack is a struct that contain a user type pointer and a verification token.
type TempVerificationPack struct {
	user  *entities.User
	token string
}

// ProjectUserContainer is a struct that containe both the project and owner information.
type ProjectUserContainer struct {
	Firstname     string
	Lastname      string
	Phonenumber   string
	Email         string
	JobTitle      string
	Country       string
	City          string
	Gender        string
	ProfilePic    string
	Project       *entities.Project
	CreatedString string
}

// IService is an interface that specifies what a User type can do.
type IService interface {
	RegisterUser(*entities.User) error
	Login(string, string) error
	EditUserProfile(*entities.User) error
	Verification(*entities.User, Identification) error
	SearchUser(string) *entities.User
	AddMatchTag(uid string, category string, subcategory string, worktype string) error
	RemoveMatchTag(uid string, category string, subcategory string, worktype string) error
	GetMatchTags(uid string) []*MatchTag
	SearchProjectWMatchTag(uid string) []*ProjectUserContainer
	// AuthUserProfile(*entities.User, Identification) error
}

// NewService is a function that returns a new UserService type.
func NewService(connection IRepository) IService {
	return &Service{conn: connection}
}

// Verification is a method that verify that eligibilty of a new user.
func (service *Service) Verification(userHolder *entities.User, idnt Identification) error {

	tempUser, origin := service.conn.SearchTPUser(userHolder.Email)

	if tempUser.UID != "" && origin == idnt.From {
		userHolder.Password = tempUser.Password
		return nil
	}
	if idnt.TpFlag {
		userHolder.Password = randomStringGN(10)
		tempUser = service.conn.SearchUser(userHolder.Email)
		if tempUser.UID != "" {
			return errors.New("email already exists")
		}
		service.RegisterUser(userHolder)
		service.conn.RegisterTPUsers(userHolder.UID, userHolder.Email, userHolder.Password, idnt.From)
		return nil
	}

	if idnt.From == "EditProfile" {
		err := service.AuthUserProfile(userHolder, idnt)
		if err != nil {
			return err
		}
		return nil
	}

	verificationToken := randomStringGN(30)
	tempVerificationPack := TempVerificationPack{user: userHolder, token: verificationToken}
	UserTempRepository = append(UserTempRepository, tempVerificationPack)
	err := service.AuthUserProfile(userHolder, idnt)
	if err != nil {
		return err
	}
	service.SendVerification(tempVerificationPack)
	return nil
}

// RegisterUser is a method that register a new user to the system.
func (service *Service) RegisterUser(user *entities.User) error {
	err := service.conn.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}

// Login is a method that validate a user using Email and Password.
func (service *Service) Login(email string, password string) error {

	user := service.conn.SearchUser(email)
	if user.UID == "" {
		return errors.New("please use a valid email addres")
	}

	if user.Password != password {
		return errors.New("incorrect password")
	}

	return nil
}

// EditUserProfile is a method that updates the user type using the provided parameter.
func (service *Service) EditUserProfile(user *entities.User) error {

	if user.ProfilePic != "" {
		filename := URepositoryDB.RemoveFileDB(user.UID, "profile_pic")
		if filename != "" {
			if err := URepositoryDB.RemoveFile(filename, "profile_pic"); err != nil {
				return err
			}
		}

	} else {
		xUser := service.SearchUser(user.UID)
		user.ProfilePic = xUser.ProfilePic
	}

	if user.CV != "" {
		filename := URepositoryDB.RemoveFileDB(user.UID, "cv")
		if filename != "" {
			if err := URepositoryDB.RemoveFile(filename, "cv"); err != nil {
				return err
			}
		}
	} else {
		xUser := service.SearchUser(user.UID)
		user.CV = xUser.CV
	}

	err := service.conn.UpdateUser(user)

	if err != nil {
		return err
	}

	tempUser, _ := service.conn.SearchTPUser(user.UID)
	if tempUser.UID != "" {
		err := service.conn.UpdateTPUsers(user.UID, user.Email)
		if err != nil {
			return err
		}
	}

	return nil
}

// SearchUser is a method that is used for searching a user from the system.
func (service *Service) SearchUser(identifier string) *entities.User {
	user := service.conn.SearchUser(identifier)
	if user.UID == "" {
		return nil
	}
	return user
}

// AddMatchTag is a method that facilitaties the adding process of match tag.
func (service *Service) AddMatchTag(uid string, category string, subcategory string, worktypes string) error {

	worktype := 0

	if !service.conn.SearchMember("categories", category) {
		return errors.New("unknown category")
	}

	if subcategory != "" && !service.conn.SearchMember("subcategories", subcategory) {
		return errors.New("unknown subcategory")
	}

	switch worktypes {
	case "Fixed":
		worktype = 1
	case "Perhour":
		worktype = 2
	case "Negotiable":
		worktype = 3
	case "":
		worktype = 4
	default:
		return errors.New("unknow work type")
	}

	matchTagStore := service.conn.GetUserMatchTags(uid)
	for _, value := range matchTagStore {
		if value.Category == category && value.Subcategory == subcategory && value.WorkType == worktype {
			return errors.New("duplicate tag")
		}
	}
	if len(matchTagStore) > 2 {
		return errors.New("reached maximum tag support")
	}

	err := service.conn.AddMatchTag(uid, category, subcategory, worktype)
	if err != nil {
		return err
	}

	return nil
}

// RemoveMatchTag is a method that facilitaties the removing of match tag of a user.
func (service *Service) RemoveMatchTag(uid string, category string, subcategory string, worktypes string) error {

	worktype := 0

	switch worktypes {
	case "Fixed":
		worktype = 1
	case "Perhour":
		worktype = 2
	case "Negotiable":
		worktype = 3
	case "":
		worktype = 4
	default:
		return errors.New("unknow work type")
	}

	if err := service.conn.RemoveMatchTag(uid, category, subcategory, worktype); err != nil {
		return err
	}

	return nil
}

// GetMatchTags is a method that returns all the match tag of a user a slice of match tags.
func (service *Service) GetMatchTags(uid string) []*MatchTag {

	matchTagStore := service.conn.GetUserMatchTags(uid)
	return matchTagStore
}

// SearchProjectWMatchTag is a method that returns all projects that match user matching tags.
func (service *Service) SearchProjectWMatchTag(uid string) []*ProjectUserContainer {

	matchTagStore := service.GetMatchTags(uid)
	var filteredProjects []*ProjectUserContainer

	for _, matchTag := range matchTagStore {

		projects := service.conn.SearchProjectWMatchTag(matchTag)

		for _, uniqueProject := range projects {
			unique := true
			for _, value := range filteredProjects {
				if uniqueProject.ID == value.Project.ID {
					unique = false
					break
				}
			}
			if unique {
				owner := service.conn.SearchUser(service.conn.GetOwner(uniqueProject.ID))
				if owner.UID != uid {
					projectUserContainer := new(ProjectUserContainer)
					projectUserContainer.Firstname = owner.Firstname
					projectUserContainer.Lastname = owner.Lastname
					projectUserContainer.Phonenumber = owner.Phonenumber
					projectUserContainer.Email = owner.Email
					projectUserContainer.JobTitle = owner.JobTitle
					projectUserContainer.Country = owner.Country
					projectUserContainer.City = owner.City
					projectUserContainer.Gender = owner.Gender
					projectUserContainer.ProfilePic = owner.ProfilePic
					projectUserContainer.Project = uniqueProject
					filteredProjects = append(filteredProjects, projectUserContainer)
				}

			}
		}
	}

	return filteredProjects
}

// AuthUserProfile is a method that authenticate a user type has a valid content or not.
func (service *Service) AuthUserProfile(user *entities.User, idnt Identification) error {

	matchFirstname, _ := regexp.MatchString(`^[a-zA-Z]\w*$`, user.Firstname)
	matchLastname, _ := regexp.MatchString(`^\w*$`, user.Lastname)
	matchEmail, _ := regexp.MatchString(`^[a-zA-Z0-9][a-zA-Z0-9\._\-&!?=#]*`, user.Email)
	matchPhonenumber, _ := regexp.MatchString(`^(\d{9})?$`, user.Phonenumber)
	matchPassword, _ := regexp.MatchString(`^\w{8}\w*$`, user.Password)
	matchGender, _ := regexp.MatchString(`^[MFO]?$`, user.Gender)

	switch {
	case !matchFirstname:
		return errors.New("invalid firstname")
	case !matchLastname:
		return errors.New("invalid lastname")
	case !matchEmail:
		return errors.New("invalid email")
	case !matchPassword:
		return errors.New("invalid password")
	case !matchPhonenumber:
		return errors.New("invalid phoneNumber")
	case user.Password != idnt.ConfirmPassword:
		return errors.New("password doesn't match")
	case !matchGender:
		return errors.New("invalid gender")
	}

	tempUser := service.conn.SearchUser(user.Email)

	if tempUser.UID != "" && idnt.From != "EditProfile" {
		return errors.New("email already exists")
	}

	return nil
}

// SendVerification is a method that take a temporary verification packet and send a verification email to provided email address.
func (service *Service) SendVerification(toBeVerified TempVerificationPack) {
	auth := smtp.PlainAuth("", "biniyam.s72@gmail.com", "0911732688", "smtp.gmail.com")
	to := []string{toBeVerified.user.Email}
	msg := []byte("To:" + toBeVerified.user.Email + "\r\n" +
		"Subject: Verify your email address\r\n" + "\r\n" +
		`Verify your email address to complete registration` + "\r\n" + "\r\n" +
		`Hi ` + toBeVerified.user.Firstname + ",\r\n" + "\r\n" +
		`Thanks for your interest in Fjobs! To complete your registration, we need you to verify your email address.` + "\r\n" +
		"Open the Link to verify: " +
		`http://localhost:1234/Verify?token=` + toBeVerified.token + "&email=" + toBeVerified.user.Email)
	err := smtp.SendMail("smtp.gmail.com:587", auth, "biniyam.s72@gmail.com", to, msg)
	if err != nil {
		panic(err)
	}
}

func randomStringGN(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
