package controllers

import (
	"net/http"

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
