package utils

import (
	"fmt"
	"net/http"
	"strings"
	c "weather-api/client"
)

func validateCityName(cit string) (string, error) {
	var city string
	switch  {
	case :
		
	}
	if strings.ContainsAny(cit, " ") {
		city = strings.Replace(cit, " ", "%20", -1)
	} else {
		city = cit
	}
}
func GeoClient(city string) (*http.Response, error) {
	url := "https://wft-geo-db.p.rapidapi.com/v1/geo/cities?limit=1&namePrefix=" + city + "&sort=-population"
	req, err := c.HttpGetRequest("GET", url, "wft-geo-db.p.rapidapi.com", nil)
	if err != nil {
		return req.Response, fmt.Errorf("error occured with httpGetRequest: %v", err)
	}
	res, err := c.HttpGetResponse(req)
	if err != nil {
		return res, fmt.Errorf("error occured with httpGetResponse: %v", err)
	}
	return res, nil
}

func GetCityLocation(cit string) (alti, longi string, e error) {
	url := "https://wft-geo-db.p.rapidapi.com/v1/geo/cities?limit=1&namePrefix=" + cit + "&sort=-population"
	req, err := c.HttpGetRequest("GET", url, "wft-geo-db.p.rapidapi.com", nil)
	if err != nil {
		return "", "", fmt.Errorf("error occured with httpGetRequest: %v", err)
	}
	res, err := c.HttpGetResponse(req)
	if err != nil {
		return "", "", fmt.Errorf("error occured with httpGetResponse: %v", err)
	}
	defer res.Body.Close()
	body, err := JsonDecoder(res.Body)
	if err != nil {
		return "", "", fmt.Errorf("cannot decode Json , error: %v", err)
	}
	convMap, err := ConvertMapLocation(body)
	if err != nil {
		return "", "", fmt.Errorf("error occured with response body: %v", err)
	}
	alti, longi, err = CityLocation(convMap)
	if err != nil {
		return "", "", fmt.Errorf("error occured during getting the value for CityLocation: %v", err)
	}
	return alti, longi, nil
}
