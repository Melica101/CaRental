package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	//"strconv"
	"CarRental/entities"
	"CarRental/form"
	"CarRental/rtoken"
	"CarRental/spam"
)

type AdminDashboardHandler struct {
	tmpl        *template.Template
	spamService spam.SpamService
	csrfSignKey []byte
}

func NewAdminDashboardHandler(
	t *template.Template,
	spamService spam.SpamService,
	csrfSignKey []byte,
) *AdminDashboardHandler {
	return &AdminDashboardHandler{tmpl: t, spamService: spamService, csrfSignKey: csrfSignKey}
}

func (adminDashboardHandler *AdminDashboardHandler) AdminIndex(w http.ResponseWriter, r *http.Request) {
	currentSession, _ := r.Context().Value(ctxUserSessionKey).(*entities.Session)
	spam, errs := adminDashboardHandler.spamService.GetSpam(currentSession.UUID)
	if len(errs) > 0 {
		http.Redirect(w, r, "/adminhome", http.StatusSeeOther)
		return
	}

	err := adminDashboardHandler.tmpl.ExecuteTemplate(w, "adminhome.layout", spam)
	fmt.Println(err)
}

func (adminDashboardHandler *AdminDashboardHandler) AdminSpams(w http.ResponseWriter, r *http.Request) {
	currentSession, _ := r.Context().Value(ctxUserSessionKey).(*entities.Session)
	spam, errs := adminDashboardHandler.spamService.GetSpam(currentSession.UUID)
	if len(errs) > 0 {
		http.Redirect(w, r, "/adminhome", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodGet {
		CSFRToken, err := rtoken.GenerateCSRFToken(adminDashboardHandler.csrfSignKey)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		spams, errs := adminDashboardHandler.spamService.GetSpam(spam.ID)

		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		spamsData := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			Spams   (*entities.Spam)
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			Spams:   spams,
			CSRF:    CSFRToken,
		}
		adminDashboardHandler.tmpl.ExecuteTemplate(w, "feedback.layout", spamsData)
	}
}
