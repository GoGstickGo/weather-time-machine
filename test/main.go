package main

import (
	"fmt"
	"log"
	"strings"
	c "weather-api/client"
	u "weather-api/utils"
)

func GetCityLocation(cit string) (alti, longi string, e error) {
	var city string
	if strings.ContainsAny(cit, " ") {
		city = strings.Replace(cit, " ", "%20", -1)
	} else {
		city = cit
	}
	url := "https://wft-geo-db.p.rapidapi.com/v1/geo/cities?limit=1&namePrefix=" + city + "&sort=-population"
	fmt.Println(url)
	req, err := c.HttpGetRequest("GET", url, "wft-geo-db.p.rapidapi.com", nil)
	if err != nil {
		return "", "", fmt.Errorf("error occured with httpGetRequest: %v", err)
	}
	res, err := c.HttpGetResponse(req)
	if err != nil {
		return "", "", fmt.Errorf("error occured with httpGetResponse: %v", err)
	}
	defer res.Body.Close()
	body, err := u.JsonDecoder(res.Body)
	if err != nil {
		return "", "", fmt.Errorf("cannot decode Json , error: %v", err)
	}
	convMap, err := u.ConvertMapLocation(body)
	if err != nil {
		return "", "", fmt.Errorf("error occured with response body: %v", err)
	}
	alti, longi, err = u.CityLocation(convMap)
	if err != nil {
		return "", "", fmt.Errorf("error occured during getting the value for CityLocation: %v", err)
	}
	return alti, longi, nil
}

func main() {
	a, b, err := GetCityLocation("Portland")
	if err != nil {
		log.Fatalf("error occured during requiring altitude ,longtitude values ,error: %v ", err)
	}
	fmt.Println(a, b)

}
