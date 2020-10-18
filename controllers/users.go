package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/meilleur/views"
)

// NewUsers is used to create a new Users controller
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// Users struct
type Users struct {
	NewView *views.View
}

// New is used for render template
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// SignupForm data for parsing form using gorilla/schema
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create is for creating new user
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	dec := schema.NewDecoder()
	var form SignupForm
	if err := dec.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, form)

	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostForm["password"])

}
