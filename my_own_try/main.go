package main

import (
	"fmt"
	"log"
	"net/http"
	"web/my_own_try/views"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var homeView, loginView *views.View
var decoder = schema.NewDecoder()

func home(wr http.ResponseWriter, r *http.Request) {
	err := homeView.Render(wr, nil)
	if err != nil {
		panic(err)
	}
}

func login(wr http.ResponseWriter, r *http.Request) {
	err := loginView.Render(wr, nil)
	if err != nil {
		panic(err)
	}
}

func loginPost(wr http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var person struct {
		Email    string `schema:email`
		Password string `schema:password`
	}

	decoder.Decode(&person, r.PostForm)

	fmt.Println(person)
	fmt.Fprintln(wr, "Hello there!")

}

func main() {

	//Making a view.
	homeView = views.New("bs", "views/index.html")
	loginView = views.New("bs", "views/login.html")

	//Router
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/login", login).Methods("GET")
	router.HandleFunc("/login", loginPost).Methods("POST")

	//Starting Web Server
	log.Println("Webserver is Running.")
	http.ListenAndServe(":8080", router)

}
