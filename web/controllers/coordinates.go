package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"weather-api/rapidapis"
	"weather-api/utils"
	"weather-api/web/models"
	"weather-api/web/views"
)

type CoordinatesRequest struct {
	CoordinatesView       *views.View
	CoordinatesReturnView *views.View
	c                     *models.CoordinatesService
}

type CoordinatesForm struct {
	Latitude   string `schema:"latitude"`
	Longitude  string `schema:"longitude"`
	Day        string `schema:"day"`
	Month      string `schema:"month"`
	Year       string `schema:"year"`
	Fahrenheit bool   `schema:"fahrenheit"`
}

func NewRequestCo(c *models.CoordinatesService) *CoordinatesRequest {
	return &CoordinatesRequest{
		CoordinatesView:       views.NewView("bootstrap", "dynamic/coordinates"),
		CoordinatesReturnView: views.NewView("bootstrap", "dynamic/coordinatesreturn"),
		c:                     c,
	}
}

func (c *CoordinatesRequest) New(w http.ResponseWriter, r *http.Request) {
	if err := c.CoordinatesView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (c *CoordinatesRequest) GetTempsCo(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form CoordinatesForm

	if err := parseForm(r, &form); err != nil {
		log.Println(err)
		vd.Alert = &views.Alert{
			Level:   views.AlertLvlError,
			Message: views.AlertGenericMsg,
		}
	}
	date, err := utils.BuildDate(form.Year, form.Month, form.Day)
	if err != nil {
		log.Println(err)
		vd.Alert = &views.Alert{
			Level:   views.AlertLvlError,
			Message: fmt.Sprintf("%v", err),
		}
	}
	freeCoordinates, _ := c.c.FreeCoordinates(date, form.Latitude, form.Longitude, form.Fahrenheit)
	if err != nil {
		log.Println(err)
		vd.Alert = &views.Alert{
			Level:   views.AlertLvlError,
			Message: fmt.Sprintf("%v", err),
		}
	}
	if freeCoordinates != nil {
		log.Println("free", freeCoordinates)
		c.CoordinatesReturnView.Render(w, freeCoordinates)
	} else {
		tempL, tempH, date, lat, long, f, err := rapidapis.DsReturnsCoWeb(rapidapis.Params{
			Latitude:   form.Latitude,
			Longitude:  form.Longitude,
			Day:        form.Day,
			Month:      form.Month,
			Year:       form.Year,
			Fahrenheit: form.Fahrenheit,
			Apikey:     os.Getenv("RAPIDAPI_KEY"),
		})
		if err != nil {
			log.Println(err)
			vd.Alert = &views.Alert{
				Level:   views.AlertLvlError,
				Message: fmt.Sprintf("%v", err),
			}
			c.CoordinatesView.Render(w, vd)
			return
		}
		coordinates := models.Coordinates{
			Latitude:   lat,
			Longitude:  long,
			Date:       date,
			TempHigh:   tempH,
			TempLow:    tempL,
			Fahrenheit: f,
		}

		if err := c.c.CreateCo(&coordinates); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Print("not free")
		c.CoordinatesReturnView.Render(w, coordinates)
	}

}
