package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate, contactTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		_panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		_panic(err)
	}
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.html", "views/layouts/footer.html")
	_panic(err)
	contactTemplate, err = template.ParseFiles("views/contact.html", "views/layouts/footer.html")
	_panic(err)

	mux := mux.NewRouter()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/contact", contact)

	http.ListenAndServe(":8080", mux)

}
