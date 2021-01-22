package main

import (
	"log"
	"net/http"
	"web/views"

	"github.com/gorilla/mux"
)

var homeView, contactView, faqView *views.View

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Render(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Render(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	if err := faqView.Render(w, nil); err != nil {
		panic(err)
	}
}
func main() {
	homeView = views.NewView("bootstrap", "views/home.html")
	contactView = views.NewView("bootstrap", "views/contact.html")
	faqView = views.NewView("bootstrap", "views/faq.html")
	mux := mux.NewRouter()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/faq", faq)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
