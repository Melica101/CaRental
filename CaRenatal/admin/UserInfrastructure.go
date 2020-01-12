package user

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/amalica/CarRental/entities"
	// "github.com/gorilla/sessions"
	// "golang.org/x/oauth2"
	// "golang.org/x/oauth2/facebook"
	// "golang.org/x/oauth2/google"
	// "golang.org/x/oauth2/linkedin"
)

const htmlIndex = `<html><body>
<a href="/GoogleLogin">Log in with Google</a>
<a href="/LinkedInLogin">Log in with linked in</a>
<a href="/FacebookLogin">Log in with Facebook</a>
</body></html>
`

var (

	// URepositoryDB is a pointer to the user PsqlUserRepository type.
	URepositoryDB IRepository
	// UService is a pointer to the user Service type.
	UService IService

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:1234/GoogleCallback",
		ClientID:     "866699224887-chqgletm3pgv85d8t1j3hqu9qbggg9c6.apps.googleusercontent.com",
		ClientSecret: "EcGwbY1ikvFKs2kPhKSQPk8X",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}

	linkedinOauthConfig = &oauth2.Config{
		ClientID:     "86uo5nbpfdo1ac",
		ClientSecret: "bJSiYAeW4XhhlaTB",
		RedirectURL:  "http://localhost:1234/LinkedInCallback",
		Scopes:       []string{"r_liteprofile", "r_emailaddress"},
		Endpoint:     linkedin.Endpoint,
	}

	facebookOauthConfig = &oauth2.Config{
		ClientID:     "453733495330808",
		ClientSecret: "4b5856af660b4440654e8c4eb31c3a45",
		RedirectURL:  "http://localhost:1234/FacebookCallback",
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	}
	// ServerAUT is a sceret key that is used to identify the POST request, if it checks out then the call was made from the server it self.
	ServerAUT string
	// OauthStateString is a sceret state key that is used for identifying a request in api exchange.
	OauthStateString string
	// UserTempRepository is a slice that contain a temporary verificaion packet that are waiting to be verified.
	UserTempRepository []TempVerificationPack

	key         = []byte("cookie-secret-key")
	cookieStore = sessions.NewCookieStore(key)

	// Temp contains all the parsed templates.
	Temp = template.Must(template.ParseGlob("templates/*.html"))
)

// HandleMain is a Handler func that parse the main window.
func HandleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

// HandleGoogleLogin is a Handler func that initiate google oauth process to the google.Endpoint.
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	OauthStateString = randomStringGN(20)
	url := googleOauthConfig.AuthCodeURL(OauthStateString)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// HandleGoogleCallback is a Handler func that handle all the token verification and profile information from google api.
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	statCheck := r.FormValue("state")
	if statCheck != OauthStateString {
		http.Error(w, fmt.Sprintf("Wrong state string: Expected %s, got %s. Please, try again", OauthStateString, statCheck), http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Code exchange failed with %s", err.Error()), http.StatusBadRequest)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	var tempUser struct {
		Sub           string `json:"sub"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Profile       string `json:"profile"`
		Picture       string `json:"picture"`
		Email         string `json:"email"`
		EmailVerified string `json:"email_verified"`
		Gender        string `json:"gender"`
	}
	err = json.Unmarshal(contents, &tempUser)
	if err != nil {
		panic(err)
	}
	ServerAUT = randomStringGN(20)
	postQuery := "/Register?thirdParty=true" +
		"&firstname=" + tempUser.GivenName +
		"&lastname=" + tempUser.FamilyName +
		"&email=" + tempUser.Email +
		"&from=google" +
		"&serverAUT=" + ServerAUT
	http.Redirect(w, r, postQuery, http.StatusSeeOther)

}

// HandleFacebookLogin is a Handler func that initiate google oauth process to the facebook.Endpoint.
func HandleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	OauthStateString = randomStringGN(20)
	url := facebookOauthConfig.AuthCodeURL(OauthStateString)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// HandleFacebookCallback is a Handler func that handle all the token verification and profile information from facebook api.
func HandleFacebookCallback(w http.ResponseWriter, r *http.Request) {

	statCheck := r.FormValue("state")
	if statCheck != OauthStateString {
		http.Error(w, fmt.Sprintf("Wrong state string: Expected %s, got %s. Please, try again", OauthStateString, statCheck), http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Code exchange failed with %s", err.Error()), http.StatusBadRequest)
		return
	}

	response, err := http.Get(`https://graph.facebook.com/me?access_token=` + token.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)

	var checkToken struct {
		ID   string `json: "id"`
		Name string `json:"name"`
	}
	err = json.Unmarshal(contents, &checkToken)
	if err != nil {
		panic(err)
	}
	response, err = http.Get(`https://graph.facebook.com/` + checkToken.ID + `?fields=id,name,email,first_name,last_name&access_token=` + token.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	contents, err = ioutil.ReadAll(response.Body)

	var tempUser struct {
		Name      string `json:"name"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
	err = json.Unmarshal(contents, &tempUser)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(tempUser.Email))

	ServerAUT = randomStringGN(20)
	postQuery := "/Register?thirdParty=true" +
		"&firstname=" + tempUser.FirstName +
		"&lastname=" + tempUser.LastName +
		"&email=" + tempUser.Email +
		"&from=facebook" +
		"&serverAUT=" + ServerAUT

	http.Redirect(w, r, postQuery, http.StatusSeeOther)

}

// HandleLinkedInLogin is a Handler func that initiate google oauth process to the linkedin.Endpoint.
func HandleLinkedInLogin(w http.ResponseWriter, r *http.Request) {
	OauthStateString = randomStringGN(20)
	url := linkedinOauthConfig.AuthCodeURL(OauthStateString)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// HandleLinkedInCallback is a Handler func that handle all the token verification and profile information from linkedin api.
func HandleLinkedInCallback(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	statCheck := r.FormValue("state")
	if OauthStateString != statCheck {
		http.Error(w, fmt.Sprintf("Wrong state string: Expected %s, got %s. Please, try again", OauthStateString, statCheck), http.StatusBadRequest)
		return
	}

	token, err := linkedinOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := linkedinOauthConfig.Client(oauth2.NoContext, token)
	req, err := http.NewRequest("GET", "https://api.linkedin.com/v2/me", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req.Header.Set("Bearer", token.AccessToken)
	response, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer response.Body.Close()
	str, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var tempUser struct {
		ID        string
		FirstName string `json:"localizedFirstName"`
		LastName  string `json:"localizedLastName"`
		Email     string `json:"emailAddress"`
		Pp        struct {
			Picture string `json:"displayImage"`
		} `json:"profilePicture"`
	}

	err = json.Unmarshal(str, &tempUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServerAUT = randomStringGN(20)
	postQuery := "/Register?thirdParty=true" +
		"&firstname=" + tempUser.FirstName +
		"&lastname=" + tempUser.LastName +
		"&email=" + tempUser.Email +
		"&from=linkedin" +
		"&serverAUT=" + ServerAUT
	http.Redirect(w, r, postQuery, http.StatusSeeOther)

}

// Register is a Handler func that initaite registration process.
func Register(w http.ResponseWriter, r *http.Request) {

	var userHolder *entities.User
	var password string

	thirdParty := r.FormValue("thirdParty")
	var identification Identification
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	identification.ConfirmPassword = r.FormValue("confirmPassword")

	if thirdParty == "true" {

		if r.FormValue("serverAUT") != ServerAUT {
			http.Error(w, "Invalid server key", http.StatusBadRequest)
			return
		}

		identification.From = r.FormValue("from")
		identification.TpFlag = true
	} else {
		password = r.FormValue("password")
		identification.ConfirmPassword = r.FormValue("confirmPassword")
	}

	userHolder = entities.NewUserFR(firstname, lastname, email, password)

	err := UService.Verification(userHolder, identification)
	if err != nil {
		panic(err)
	}
	if identification.TpFlag {
		session, _ := cookieStore.Get(r, "Fjobs_User_Cookie")
		session.Values["email"] = email
		session.Values["auth"] = true
		session.Save(r, w)

		http.Redirect(w, r, "/Dashboard", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/Check_Your_Email", http.StatusSeeOther)
	}
	return

}

// Login is a Handler func that initaite login process.
func Login(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	err := UService.Login(email, password)

	if err != nil {
		panic(err)
	}

	session, _ := cookieStore.Get(r, "Fjobs_User_Cookie")
	session.Values["email"] = email
	session.Values["auth"] = true
	session.Save(r, w)

	http.Redirect(w, r, "/Dashboard", http.StatusSeeOther)
}

// UpdateProfile is a Handler func that initaite the updating profile process.
func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	loggedInUser := Authentication(r)
	if loggedInUser == nil {
		http.Redirect(w, r, "/Login_Page", http.StatusSeeOther)
		return
	}

	var profilePic, cv string
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	phonenumber := r.FormValue("phonenumber")
	email := r.FormValue("email")
	jobTitle := r.FormValue("jobTitle")
	country := r.FormValue("country")
	city := r.FormValue("city")
	gender := r.FormValue("gender")
	bio := r.FormValue("bio")
	// In Edit profile password will not be change be still it will be authenticated!
	password := "ValidPassword123"

	user := entities.NewUser(firstname, lastname, email, profilePic, password, phonenumber, jobTitle, country, city, gender, cv, bio)
	user.UID = loggedInUser.UID

	err := UService.Verification(user, Identification{ConfirmPassword: "ValidPassword123", From: "EditProfile"})
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}
	user.ProfilePic = ResourceExtractor(user.UID, "profilePic", "image", r)
	user.CV = ResourceExtractor(user.UID, "cv", "file", r)
	err = UService.EditUserProfile(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	w.Write([]byte("Okay"))
}

// Dashboard is a Handler func that initaite the Home page of a user after checking his/her profile is completed.
func Dashboard(w http.ResponseWriter, r *http.Request) {

	user := Authentication(r)
	if user == nil {
		http.Redirect(w, r, "/Login_Page", http.StatusSeeOther)
		return
	}

	// Means incomplete profile.
	if user.Gender == "" {
		http.Redirect(w, r, "/EditProfile", http.StatusSeeOther)
		return
	}

	Temp.ExecuteTemplate(w, "Dashboard.html", user)

}

// Verify is a Handler func that verify a token from a request query and user to database if valid.
func Verify(w http.ResponseWriter, r *http.Request) {
	verificationToken := r.URL.Query().Get("token")
	email := r.URL.Query().Get("email")

	for _, value := range UserTempRepository {
		if value.token == verificationToken && value.user.Email == email {
			err := UService.RegisterUser(value.user)
			if err != nil {
				panic(err)
			}
			session, _ := cookieStore.Get(r, "Fjobs_User_Cookie")
			session.Values["email"] = email
			session.Values["auth"] = true
			session.Save(r, w)
			http.Redirect(w, r, "/Dashboard", http.StatusSeeOther)
			return
		}
	}
	http.Error(w, "invalid verification Token!", http.StatusBadRequest)
}

// EditProfile is a Handler func that initaite the Editing profile process.
func EditProfile(w http.ResponseWriter, r *http.Request) {

	user := Authentication(r)
	if user == nil {
		http.Redirect(w, r, "/Login_Page", http.StatusSeeOther)
		return
	}

	Temp.ExecuteTemplate(w, "ProfilePage.html", user)
}

// Authentication is a function that checks if the user has already logged in or not.
func Authentication(r *http.Request) *entities.User {

	session, _ := cookieStore.Get(r, "Fjobs_User_Cookie")
	email, ok := session.Values["email"].(string)

	if !ok {
		return nil
	}
	auth, ok := session.Values["auth"].(bool)
	if !ok || !auth {
		return nil
	}

	user := UService.SearchUser(email)
	return user
}

// Logout is a Handler func that perform logout operation by revoking authentication pass.
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := cookieStore.Get(r, "Fjobs_User_Cookie")

	session.Values["auth"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/Login_Page", http.StatusSeeOther)
}

// AddMatchTag is a Handler func that accepts the matching tags store them.
func AddMatchTag(w http.ResponseWriter, r *http.Request) {

	user := Authentication(r)
	if user == nil {
		http.Redirect(w, r, "/Login_Page", http.StatusSeeOther)
		return
	}
	uid := user.UID
	category := r.FormValue("category")
	subcategory := r.FormValue("subcategory")
	worktype := r.FormValue("worktype")

	if err := UService.AddMatchTag(uid, category, subcategory, worktype); err != nil {
		panic(err)
	}
	w.Write([]byte("okay"))

}

// RemoveMatchTag is a Handler func that accepts the matching tags and remove it.
func RemoveMatchTag(w http.ResponseWriter, r *http.Request) {

	user := Authentication(r)
	if user == nil {
		http.Redirect(w, r, "/Login_Page", http.StatusSeeOther)
		return
	}
	uid := user.UID
	category := r.FormValue("category")
	subcategory := r.FormValue("subcategory")
	worktype := r.FormValue("worktype")

	if err := UService.RemoveMatchTag(uid, category, subcategory, worktype); err != nil {
		panic(err)
	}

	w.Write([]byte("okay"))
}

// GetMatchTags is a Handler func that sends all the match tag the user have.
func GetMatchTags(w http.ResponseWriter, r *http.Request) {
	user := Authentication(r)
	if user == nil {
		http.Redirect(w, r, "/Login_Page", http.StatusSeeOther)
		return
	}

	matchTagStore := UService.GetMatchTags(user.UID)
	json, err := json.Marshal(matchTagStore)
	if err != nil {
		panic(err)
	}

	w.Write(json)

}

// GetProjectsWMatchTags is a Handler func that sends all the projects taht match the user match Tags.
func GetProjectsWMatchTags(w http.ResponseWriter, r *http.Request) {

	user := Authentication(r)
	if user == nil {
		http.Redirect(w, r, "/Login_Page", http.StatusSeeOther)
		return
	}

	uid := user.UID

	projects := UService.SearchProjectWMatchTag(uid)

	json, err := json.Marshal(projects)
	if err != nil {
		panic(err)
	}

	w.Write(json)
}

// ResourceExtractor is a method that extract file from a request.
func ResourceExtractor(uid string, name string, fileType string, r *http.Request) string {

	count := randomStringGN(10)

	fm, fh, err := r.FormFile(name)
	if err != nil {
		return ""
	}
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

	NewFileName := fmt.Sprintf("asset_"+fileType+"_"+uid+"_%s"+suffix, count)
	if name == "profilePic" {
		column := "profile_pic"
		path = filepath.Join(path, "assets", column, NewFileName)
	} else {
		path = filepath.Join(path, "assets", name, NewFileName)
	}

	out, _ := os.Create(path)
	defer out.Close()

	_, err = io.Copy(out, fm)

	if err != nil {
		panic(err)
	}

	return NewFileName

}

// Index is check function.
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sucessfull!"))
}

// IndexNot is check function.
func IndexNot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Sucessfull!"))
}
