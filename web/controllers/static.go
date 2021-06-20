package controllers

import "weather-api/web/views"

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
		History: views.NewView("bootstrap", "static/history"),
	}
}

type Static struct {
	Home    *views.View
	Contact *views.View
	History *views.View
}
