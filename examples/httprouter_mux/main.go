package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func homePage(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Fprintln(wr, "<h1>Welcome to homepage</h1>")

}

func main() {
	router := httprouter.New()
	router.GET("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", router))
}
