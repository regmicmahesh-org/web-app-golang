package controllers

import (
	"fmt"
	"net/http"
	"web/views"
)

//Users :This is user Controller
type Users struct {
	NewView *views.View
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bs", "views/users/new.html"),
	}
}

//GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

//POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, r.PostForm.Get("email"))
	fmt.Fprintln(w, r.PostForm.Get("password"))
}
