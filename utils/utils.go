package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"time"
	"weather-api/defaults"
)

var (
	checkY = regexp.MustCompile(`\b(19[4-9][0-9]|20[0-4][0-9]|2050)\b`).MatchString
	checkD = regexp.MustCompile(`^(19|20)\d\d[- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01])$`).MatchString
)

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

func BuildBaseURL(latitude, longitude, date string) string {
	baseurl := defaults.DarkSkyApiUrl + latitude + "," + longitude + "," + date + "T" + "12:00:00" + defaults.DarkSkyApiSort
	return baseurl

}
func JsonDecoder(r io.Reader) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	return data, json.NewDecoder(r).Decode(&data)
}
