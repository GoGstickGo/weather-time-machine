package geodb

import (
	"fmt"
	"log"
	"weather-api/client"
	"weather-api/defaults"
)

var (
	City = "Dublin"
	c    = client.Responses{
		Url:        defaults.GeoDBUrl + City + defaults.GeoDBUrlSort,
		Method:     defaults.GET,
		Apiaddress: defaults.GeoDBApi,
	}
	m = GeoDB{
		MapNameGeoDB: defaults.GeoDBMap,
	}
)

func GeoDBreturns(m *GeoDB) (string, string, string, error) {
	c.Data, c.Error = client.Client(&c)
	if c.Error != nil {
		log.Fatalf("error occured with GeoDB client: %v", c.Error)
	}
	m.Recieved = c.Data
	m.TempField, m.Error = convertMap(m)
	if m.Error != nil {
		return "", "", "", fmt.Errorf("error occured getting values from GeoDB: %v", m.Error)
	}
	m.Latitude, m.Longitude, m.Error = getCityLocation(m)
	if m.Error != nil {
		return "", "", "", fmt.Errorf("error occured when getting cordinates for the %s, error:%v", City, m.Error)
	}
	m.CountryCode, m.Error = getCountryCode(m)
	if m.Error != nil {
		return "", "", "", fmt.Errorf("error occured when getting countryCode for the %s, error:%v", City, m.Error)
	}
	return m.Latitude, m.Longitude, m.CountryCode, nil
}
