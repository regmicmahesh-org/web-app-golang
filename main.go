package main

import (
	"net/http"
	"web/controllers"
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
	usersC := controllers.NewUsers()
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/signup", usersC.New).Methods("GET")
	router.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":8080", router)
}
