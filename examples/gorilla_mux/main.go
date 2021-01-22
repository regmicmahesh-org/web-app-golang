package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello there!")
}

func faq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
<h1> Here are all the frequently asked questions. </h1>	
<h2> oops no questions </h2>
`)
}

type Handler struct{}

func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Failed to load the page :( sorry to hear bro.")
}

func main() {
	var handler = Handler{}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = handler
	fmt.Println("Currently running..")
	http.ListenAndServe(":8080", r)
}
