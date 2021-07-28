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

type CityRequest struct {
	CityView       *views.View
	CityReturnView *views.View
	c              *models.CityService
}

type CityForm struct {
	Name       string `schema:"name"`
	Day        string `schema:"day"`
	Month      string `schema:"month"`
	Year       string `schema:"year"`
	Fahrenheit bool   `schema:"fahrenheit"`
}

func NewRequest(c *models.CityService) *CityRequest {
	return &CityRequest{
		CityView:       views.NewView("bootstrap", "dynamic/city"),
		CityReturnView: views.NewView("bootstrap", "dynamic/cityreturn"),
		c:              c,
	}
}

func (c *CityRequest) New(w http.ResponseWriter, r *http.Request) {
	if err := c.CityView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (c *CityRequest) GetTemps(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form CityForm

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
	freeCities, _ := c.c.FreeCity(date, form.Name, form.Fahrenheit)
	if err != nil {
		log.Println(err)
		vd.Alert = &views.Alert{
			Level:   views.AlertLvlError,
			Message: fmt.Sprintf("%v", err),
		}
	}
	if freeCities != nil {
		log.Println("free", freeCities)
		c.CityReturnView.Render(w, freeCities)
	} else {
		tempH, tempL, date, city, cc, f, err := rapidapis.DsReturnsWeb(rapidapis.Params{
			City:       form.Name,
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
			c.CityView.Render(w, vd)
			return
		}
		cities := models.Cities{
			City:        city,
			Date:        date,
			CountryCode: cc,
			TempHigh:    tempH,
			TempLow:     tempL,
			Fahrenheit:  f,
		}

		if err := c.c.Create(&cities); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Print("not free")
		c.CityReturnView.Render(w, cities)
	}

}
