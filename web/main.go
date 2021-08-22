package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"weather-api/web/controllers"
	"weather-api/web/models"

	"github.com/gorilla/mux"
)

var (
	host     = os.Getenv("DATABASE_URL")
	port, _  = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	user     = os.Getenv("DATABASE_USER")
	password = os.Getenv("DATABASE_PASSWORD")
	dbname   = os.Getenv("DATABASE_NAME")
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	staticC := controllers.NewStatic()

	c, err := models.NewCityService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer c.CloseCon()
	c.AutoMigrate()
	cityC := controllers.NewRequest(c)

	co, err := models.NewCoordinatesService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer co.CloseConCo()
	co.AutoMigrateCo()
	coordinatesC := controllers.NewRequestCo(co)

	r := mux.NewRouter()
	r.HandleFunc("/", staticC.Home.ServeHTTP).Methods("GET")
	r.HandleFunc("/city", cityC.New).Methods("GET")
	r.HandleFunc("/cityreturn", cityC.GetTemps).Methods("POST")
	r.HandleFunc("/coordinates", coordinatesC.New).Methods("GET")
	r.HandleFunc("/coordinatesreturn", coordinatesC.GetTempsCo).Methods("POST")
	r.HandleFunc("/contact", staticC.Contact.ServeHTTP).Methods("GET")
	http.ListenAndServe(":3000", r)

}
