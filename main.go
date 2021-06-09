package main

import (
	"fmt"
	"log"
	"weather-api/client"
	"weather-api/defaults"
	"weather-api/geodb"
	"weather-api/utils"
)

func main() {

	g := geodb.GeoDB{
		MapNameGeoDB: defaults.GeoDBMap,
	}
	Latitude, Longitude, CountryCode, error := g.GeoDBreturns()
	if error != nil {
		log.Fatalf("error occured with GeoDB:%v", error)
	}
	d := utils.DateBuild{
		Year:  defaults.TestYear,
		Month: defaults.TestMonth,
		Day:   defaults.TestDay,
	}
	d.Date, d.Error = d.BuildDate()
	if d.Error != nil {
		log.Fatalf("error occured with date: %v", d.Error)
	}
	c := client.Responses{
		Url:        utils.BaseUrl(defaults.DarkSkyApiUrl, Latitude, Longitude, d.Date, defaults.DarkSkyApiSort),
		Method:     defaults.GET,
		Apiaddress: defaults.DarkSkyApi,
	}
	c.Data, c.Error = c.Client()
	if c.Error != nil {
		log.Fatalf("error occured when gettting http response body from the client: %v", c.Error)
	}
	m := utils.Mapping{
		Recieved: c.Data,
		MapName:  defaults.DarkSkyMap,
	}
	m.TempField, m.Error = m.ConvertMap()
	if m.Error != nil {
		log.Fatalf("error occured when getting response body in DarkSkyApi:%v", m.Error)
	}
	m.HighTemp, m.Error = m.GetTempH()
	if m.Error != nil {
		log.Fatalf("error occured with HighTemp,error:%v", m.Error)
	}
	m.LowTemp, m.Error = m.GetTempL()
	if m.Error != nil {
		log.Fatalf("error occured with LowTemp,error:%v", m.Error)
	}

	fmt.Printf("Highest daily temperature was %s Celcius in %s in %s, %s\n", m.HighTemp, d.Date, geodb.City, CountryCode)
	fmt.Printf("Lowest daily temperature was %s Celcius in %s in %s, %s\n", m.LowTemp, d.Date, geodb.City, CountryCode)
}
