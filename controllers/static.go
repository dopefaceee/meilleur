package controllers

import "github.com/meilleur/views"

// NewStatic abstarction
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

// Static for layout view
type Static struct {
	Home    *views.View
	Contact *views.View
}
