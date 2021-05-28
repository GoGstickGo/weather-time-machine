package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func JsonDecoder(r io.Reader) (map[string]interface{}, error) {
	var data map[string]interface{}
	return data, json.NewDecoder(r).Decode(&data)
}

func ConvertMapTemp(b map[string]interface{}) (map[string]string, error) {
	daily := make(map[string]string)
	for k, v := range b {
		if k == "daily" {
			sv := fmt.Sprintf("%v", v)
			daily[k] = sv
		}
	}
	return daily, nil
}

func ConvertMapLocation(b map[string]interface{}) (map[string]string, error) {
	daily := make(map[string]string)
	for k, v := range b {
		if k == "data" {
			sv := fmt.Sprintf("%v", v)
			daily[k] = sv
		}
	}
	return daily, nil
}

func GetHighTemp(d map[string]string) (string, error) {
	var highTemp string
	for _, v := range d {
		s := strings.Fields(v)
		for _, v := range s {
			if strings.HasPrefix(v, "temperatureHigh:") {
				temp := strings.SplitN(v, ":", 2)
				highTemp = temp[1]
				return highTemp, nil
			} else if strings.HasPrefix(v, "temperatureMax:") {
				tempy := strings.SplitN(v, ":", 2)
				highTemp = tempy[1]
			}
			//fmt.Printf("key %v,value %v\n", k, v)
		}
	}
	return highTemp, nil
}
func GetLowTemp(d map[string]string) (string, error) {
	var lowTemp string
	for _, v := range d {
		s := strings.Fields(v)
		for _, v := range s {
			if strings.HasPrefix(v, "temperatureLow:") {
				temp := strings.SplitN(v, ":", 2)
				lowTemp = temp[1]
				return lowTemp, nil
			} else if strings.HasPrefix(v, "temperatureMin:") {
				tempy := strings.SplitN(v, ":", 2)
				lowTemp = tempy[1]

			}
			//fmt.Printf("key %v,value %v\n", k, v)
		}
	}
	return lowTemp, nil
}

func CityLocation(d map[string]string) (lalt, longi string, e error) {
	for _, v := range d {
		s := strings.Fields(v)
		for _, v := range s {
			switch {
			case strings.HasPrefix(v, "latitude:"):
				latitude := strings.SplitN(v, ":", 2)
				lalt = latitude[1]
			case strings.HasPrefix(v, "longitude:"):
				longitude := strings.SplitN(v, ":", 2)
				longi = longitude[1]
			}
		}
	}
	return lalt, longi, nil
}

func checkYear(y string) bool {
	var checkY = regexp.MustCompile(`\b(19[4-9][0-9]|20[0-4][0-9]|2050)\b`).MatchString
	return checkY(y)

}
func checkDate(y, m, d string) bool {
	dat := []string{y, m, d}
	date := strings.Join(dat, "-")
	var checkD = regexp.MustCompile(`^(19|20)\d\d[- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01])$`).MatchString
	return checkD(date)
}

func checkDigits(y, m, d string) bool {
	j1 := []string{y, m, d}
	j2 := strings.Join(j1, "-")
	var checkDig = strings.ContainsAny(j2, "abcdefghijklmnopqrstuyxz")
	return !checkDig
}

func BuildTimeStr(y, m, d string) (string, error) {
	yint, err := strconv.Atoi(y)
	if err != nil {
		return "", fmt.Errorf("Error:%v", err)
	}
	switch {
	case !checkDigits(y, m, d):
		return "", fmt.Errorf("unexpected char in date format")
	case !checkYear(y) || yint > time.Now().Year():
		return "", fmt.Errorf("invalid value for year: %s,year must between 1940-%d", y, time.Now().Year())
	case !checkDate(y, m, d):
		return "", fmt.Errorf("date format must be: 1938-02-02,invalid value for date: %s-%s-%s", y, m, d)
	}
	weatherTime := y + "-" + m + "-" + d + "T" + "12:00:00"
	return weatherTime, nil
}

func PrintDate(d string) (pd string, e error) {
	printDate := strings.TrimSuffix(d, "T12:00:00")
	return printDate, nil
}
