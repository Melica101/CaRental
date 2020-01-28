package main

import (
	"html/template"
	"net/http"
)


var Temp = template.Must(template.ParseGlob("ui/*.html"))

func CheckTemplate(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "index.html", nil)
}

func Post(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "post.html", nil)
}

func Gallery(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "Gallery.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "login.html", nil)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "signup.html", nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "whoweare.html", nil)
}

func Rent(w http.ResponseWriter, r *http.Request) {
	Temp.ExecuteTemplate(w, "rent.html", nil)
}

func main() {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("ui"))
	mux.Handle("/ui/", http.StripPrefix("/ui/", fileServer))

	mux.HandleFunc("/", CheckTemplate)
	mux.HandleFunc("/post", Post)
	mux.HandleFunc("/gallery", Gallery)
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/signup", Signup)
	mux.HandleFunc("/about", About)
	mux.HandleFunc("/rent", Rent)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
