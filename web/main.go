package main

import (
	"fmt"
	"net/http"
	"weather-api/web/controllers"
	"weather-api/web/models"

	"github.com/gorilla/mux"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	//password = "FatChanceIputPasswordHere"
	dbname = "wtm"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host, port, user, dbname)
	/*psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"*password=%s dbname=%s sslmode=disable", host, port, user password,, dbname)*/

	staticC := controllers.NewStatic()

	c, err := models.NewCityService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer c.CloseCon()
	c.AutoMigrate()
	cityC := controllers.NewRequest(c)

	r := mux.NewRouter()
	r.HandleFunc("/", staticC.Home.ServeHTTP).Methods("GET")
	r.HandleFunc("/contact", staticC.Contact.ServeHTTP).Methods("GET")
	r.HandleFunc("/history", staticC.History.ServeHTTP).Methods("GET")
	r.HandleFunc("/city", cityC.New).Methods("GET")
	r.HandleFunc("/city", cityC.GetTemps).Methods("POST")
	http.ListenAndServe(":3000", r)

}
