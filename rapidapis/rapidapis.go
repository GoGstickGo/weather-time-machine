package rapidapis

import (
	"fmt"
	"io"
	"strings"
	"weather-api/defaults"
)

type Params struct {
	Day        string
	Month      string
	Year       string
	Apikey     string
	City       string
	Latitude   string
	Longitude  string
	Fahrenheit bool
	Writer     io.Writer
}

type Mapping struct {
	recieved    map[string]interface{}
	pass        map[string]string
	value       string
	err         error
	highTemp    string
	lowTemp     string
	tempField   []string
	tempSplit   []string
	latitude    string
	longitude   string
	countryCode string
}

type GeoDBClient struct {
	data    map[string]interface{}
	err     error
	mapping Mapping
	params  Params
}

type DarkSkyClient struct {
	data        map[string]interface{}
	date        string
	err         error
	mapping     Mapping
	countryCode string
}

func (c *GeoDBClient) validateCity() error {
	c.mapping.pass = make(map[string]string)
	for k, v := range c.mapping.recieved {
		if k == "metadata" {
			c.mapping.value = fmt.Sprintf("%v", v)
			c.mapping.pass[k] = c.mapping.value
		}
		for _, v := range c.mapping.pass {
			c.mapping.tempField = strings.Fields(v)
		}
		for _, j := range c.mapping.tempField {
			if strings.Contains(j, "totalCount:0") {
				return fmt.Errorf("please choose an existing city")
			}
		}
	}
	return nil
}

func (c *GeoDBClient) convertMap() ([]string, error) {
	c.mapping.pass = make(map[string]string)
	for k, v := range c.mapping.recieved {
		if k == defaults.GeoDBMap {
			c.mapping.value = fmt.Sprintf("%v", v)
			c.mapping.pass[k] = c.mapping.value
		}
	}
	for _, v := range c.mapping.pass {
		c.mapping.tempField = strings.Fields(v)
	}
	if c.mapping.tempField == nil {
		return c.mapping.tempField, fmt.Errorf("failed to convert the map to slice")
	}
	return c.mapping.tempField, nil
}

func (c *GeoDBClient) getCityLocation() (string, string, error) {
	for _, v := range c.mapping.tempField {
		switch {
		case strings.HasPrefix(v, "latitude:"):
			c.mapping.tempSplit = strings.SplitN(v, ":", 2)
			c.mapping.latitude = c.mapping.tempSplit[1]
		case strings.HasPrefix(v, "longitude:"):
			c.mapping.tempSplit = strings.SplitN(v, ":", 2)
			c.mapping.longitude = c.mapping.tempSplit[1]
		}
	}
	if c.mapping.latitude == "" || c.mapping.longitude == "" {
		return c.mapping.latitude, c.mapping.longitude, fmt.Errorf("failed to get city coordinates")
	}
	return c.mapping.latitude, c.mapping.longitude, nil
}

func (c *GeoDBClient) getCountryCode() (string, error) {
	for _, v := range c.mapping.tempField {
		if strings.HasPrefix(v, "countryCode:") {
			c.mapping.tempSplit = strings.SplitN(v, ":", 2)
			c.mapping.countryCode = c.mapping.tempSplit[1]
		}
	}
	if c.mapping.countryCode == "" {
		return c.mapping.countryCode, fmt.Errorf("failed to get countrycode")
	}
	return c.mapping.countryCode, nil
}

func (c *DarkSkyClient) convertMap() ([]string, error) {
	c.mapping.pass = make(map[string]string)
	for k, v := range c.mapping.recieved {
		if k == defaults.DarkSkyMap {
			c.mapping.value = fmt.Sprintf("%v", v)
			c.mapping.pass[k] = c.mapping.value
		}
	}
	for _, v := range c.mapping.pass {
		c.mapping.tempField = strings.Fields(v)
	}
	if c.mapping.tempField == nil {
		return c.mapping.tempField, fmt.Errorf("unfortunately there is no historic weather data")
	}
	return c.mapping.tempField, nil
}

func (c *DarkSkyClient) getTempH() (string, error) {
	for _, v := range c.mapping.tempField {
		if strings.HasPrefix(v, "temperatureHigh:") || strings.HasPrefix(v, "temperatureMax:") {
			c.mapping.tempSplit = strings.SplitN(v, ":", 2)
			c.mapping.highTemp = c.mapping.tempSplit[1]
		}
	}
	if c.mapping.highTemp == "" {
		return c.mapping.highTemp, fmt.Errorf("failed to get highest temperature")
	}
	return c.mapping.highTemp, nil
}
func (c *DarkSkyClient) getTempL() (string, error) {
	for _, v := range c.mapping.tempField {
		if strings.HasPrefix(v, "temperatureLow:") || strings.HasPrefix(v, "temperatureMin:") {
			c.mapping.tempSplit = strings.SplitN(v, ":", 2)
			c.mapping.lowTemp = c.mapping.tempSplit[1]
		}
	}
	if c.mapping.lowTemp == "" {
		return c.mapping.lowTemp, fmt.Errorf("failed to get lowest temperature")
	}
	return c.mapping.lowTemp, nil
}
