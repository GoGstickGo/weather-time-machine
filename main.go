package main

import (
	"fmt"
	"log"
	c "weather-api/client"
	u "weather-api/utils"
)

func main() {

	y := "2018"
	m := "07"
	d := "27"
	city := ""

	dataTime, err := u.BuildTimeStr(y, m, d)
	if err != nil {
		log.Fatalf("error occured with date in the API request: %v", err)
	}

	alti, longi, err := u.GetCityLocation(city)
	if err != nil {
		log.Fatalf("error occured with city corditnates %v", err)
	}
	fmt.Println(alti, longi)
	url := "https://dark-sky.p.rapidapi.com/" + alti + "," + longi + "," + dataTime + "?units=ca&exclude=currently%2%20Cminutely%2%20Chourly%2%20Calerts%2%20Cflags"
	req, err := c.HttpGetRequest("GET", url, "dark-sky.p.rapidapi.com", nil)
	if err != nil {
		log.Fatalf("error occured with httpGetRequest: %v", err)
	}

	res, err := c.HttpGetResponse(req)
	if err != nil {
		log.Fatalf("error occured with httpGetResponse: %v", err)
	}

	defer res.Body.Close()

	body, err := u.JsonDecoder(res.Body)
	if err != nil {
		log.Fatalf("cannot decode Json , error: %v", err)
	}

	convMap, err := u.ConvertMapTemp(body)
	if err != nil {
		log.Fatalf("error occured with response body: %v", err)
	}

	printDate, err := u.PrintDate(dataTime)
	if err != nil {
		fmt.Printf("Something wrong with date in the output.")
	}

	highTemp, err := u.GetHighTemp(convMap)
	if err != nil {
		log.Fatalf("error occured during getting the value for tempareture: %v", err)
	}

	fmt.Printf("%s highest temperature of the day was %s Celsius at %s\n", city, highTemp, printDate)
	lowTemp, err := u.GetLowTemp(convMap)
	fmt.Println(lowTemp)
	if lowTemp == "" || highTemp == "" {
		log.Fatalf("there is no data for date : %s, please choose another date", printDate)
	}

	if err != nil {
		log.Fatalf("error occured during getting the value for tempareture: %v", err)

	}

	fmt.Printf("%s lowest temperature of the day was %s Celsius at %s\n", city, lowTemp, printDate)

}
