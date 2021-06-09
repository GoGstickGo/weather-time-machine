package geodb

import (
	"fmt"
	"log"
	"weather-api/client"
	"weather-api/defaults"
)

var (
	c = client.Responses{
		Url:        defaults.GeoDBUrl + m.City + defaults.GeoDBUrlSort,
		Method:     defaults.GET,
		Apiaddress: defaults.GeoDBApi,
	}
	m = GeoDB{
		MapNameGeoDB: defaults.GeoDBMap,
	}
)

func (m *GeoDB) GeoDBreturns() (string, string, string, error) {
	c.Data, c.Error = c.Client()
	if c.Error != nil {
		log.Fatalf("error occured with GeoDB client: %v", c.Error)
	}
	m.Recieved = c.Data
	m.TempField, m.Error = m.convertMap()
	if m.Error != nil {
		return "", "", "", fmt.Errorf("error occured getting values from GeoDB: %v", m.Error)
	}
	m.Latitude, m.Longitude, m.Error = m.getCityLocation()
	if m.Error != nil {
		return "", "", "", fmt.Errorf("error occured when getting cordinates for the %s, error:%v", m.City, m.Error)
	}
	m.CountryCode, m.Error = m.getCountryCode()
	if m.Error != nil {
		return "", "", "", fmt.Errorf("error occured when getting countryCode for the %s, error:%v", m.City, m.Error)
	}
	return m.Latitude, m.Longitude, m.CountryCode, nil
}
