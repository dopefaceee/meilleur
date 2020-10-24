package controllers

import (
	"fmt"
	"net/http"

	"github.com/meilleur/views"
)

// NewUsers is used to create a new Users controller
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
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
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
