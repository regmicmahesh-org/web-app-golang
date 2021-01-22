package main

import (
	"net/http"
	"web/views"

	"github.com/gorilla/mux"
)

var homeView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	err := homeView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bs", "views/home.html")
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	http.ListenAndServe(":8080", router)
}
