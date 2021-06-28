package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
	"weather-api/defaults"
)

var (
	checkY     = regexp.MustCompile(`\b(19[4-9][0-9]|20[0-4][0-9]|2050)\b`).MatchString
	checkD     = regexp.MustCompile(`^(19|20)\d\d[- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01])$`).MatchString
	checkC     = regexp.MustCompile(`^\p{Lu}.*\D$`).MatchString
	checkCoord = regexp.MustCompile(`^-?[0-9]{1,3}(?:\.[0-9]{1,10})?$`).MatchString
)

func ValidateRapidApiKey(data map[string]interface{}) error {
	for k := range data {
		if k == "message" {
			value := data["message"].(string)
			if strings.Contains(value, "You are not subscribed to this API") {
				return fmt.Errorf("you are not subscribed to this API")
			}
		}
	}
	return nil
}

func ValidateParamsCity(city string) error {
	if !checkC(city) {
		return fmt.Errorf("invalid char in city parameter")
	}
	return nil
}

func ValidateParamsApikey(apikey string) error {
	if len(apikey) != 50 {
		return fmt.Errorf("please use valid API key as parameter")
	}
	return nil
}

func ValidateCoordinates(latitude, longitude string) error {
	if !checkCoord(latitude) || !checkCoord(longitude) {
		return fmt.Errorf("invalid coordinates")
	}
	return nil
}

func BuildDate(year, month, day string) (string, error) {
	var date string
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return "", fmt.Errorf("Error:%v", err)
	}
	switch {
	case !checkY(year) || yearInt > time.Now().Year():
		return year, fmt.Errorf("invalid value for year: %s,year must between 1940-%d", year, time.Now().Year())
	case !checkD(year + "-" + month + "-" + day):
		return year + "-" + month + "-" + day, fmt.Errorf("date format must be: 1978-02-02")
	default:
		date = year + "-" + month + "-" + day
	}
	return date, nil
}

func DarkSkyBuildBaseURL(latitude, longitude, date string) string {
	baseurl := defaults.DarkSkyApiUrl + latitude + "," + longitude + "," + date + "T" + "12:00:00" + defaults.DarkSkyApiSort
	return baseurl
}

func GeoDBBuildBaseURL(city string) string {
	cit := strings.Replace(city, " ", "%20", -1)
	baseurl := defaults.GeoDBUrl + cit + defaults.GeoDBUrlSort
	return baseurl
}
func JsonDecoder(r io.Reader) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	return data, json.NewDecoder(r).Decode(&data)
}
