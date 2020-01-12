package main

import (
	"html/template"
	"net/http"
)

var Temp = template.Must(template.ParseGlob("templates/*.html"))

func CheckTemplate(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "adminhome.html", nil)
}

func Post(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "post.html", nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "addpayment.html", nil)
}

func Admin(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "adminhome.html", nil)
}

func ManageV(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "manageV.html", nil)
}

func ManageR(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "manageR.html", nil)
}

func Feed(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "feedback.html", nil)
}

func Test(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "testimonials.html", nil)
}

func main() {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("templates"))
	mux.Handle("/templates/", http.StripPrefix("/templates/", fileServer))

	mux.HandleFunc("/", CheckTemplate)
	mux.HandleFunc("/post", Post)
	mux.HandleFunc("/add", Add)
	mux.HandleFunc("/admin", Admin)
	mux.HandleFunc("/manageV", ManageV)
	mux.HandleFunc("/manageR", ManageR)
	mux.HandleFunc("/feed", Feed)
	mux.HandleFunc("/testimonials", Test)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
