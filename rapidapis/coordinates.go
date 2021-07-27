package rapidapis

import (
	"fmt"
	"log"
	"weather-api/clients"
	"weather-api/defaults"
	"weather-api/utils"
)

func (p *Params) validateParamsCo() (string, error) {
	err := utils.ValidateCoordinates(p.Latitude, p.Longitude)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	err = utils.ValidateParamsApikey(p.Apikey)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	date, err := utils.BuildDate(p.Year, p.Month, p.Day)
	if err != nil {
		return "", fmt.Errorf("invalid date form %s,%v", date, err)
	}
	return date, nil
}

func dsClientCo(p Params) (*DarkSkyClient, error) {
	var url string

	date, err := (&p).validateParamsCo()
	if err != nil {
		log.Fatalf("❌ Parameter validation failed: %v", err)
	}

	switch p.Fahrenheit {
	case true:
		url = utils.DarkSkyBuildBaseURLFahrenheit(p.Latitude, p.Longitude, date)
	default:
		url = utils.DarkSkyBuildBaseURLCelcius(p.Latitude, p.Longitude, date)
	}

	response, err := clients.CreateClient(p.Apikey, defaults.GET, defaults.DarkSkyApi, url)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer response.Body.Close()

	data, err := utils.JsonDecoder(response.Body)
	if err != nil {
		return nil, fmt.Errorf("when decoding http.Response.Body to json")
	}
	return &DarkSkyClient{
		data: data,
		err:  err,
		date: date,
	}, nil
}

func DsReturnsCo(p Params) error {
	c, err := dsClientCo(p)
	if err != nil {
		log.Fatalf("❌ Error occured establishing client for DarkSky Api %v", err)
	}
	c.mapping.recieved = c.data
	c.mapping.tempField, c.mapping.err = c.convertMap()
	if c.mapping.err != nil {
		log.Fatalf("❌ DarkSkyApi error: %v, please choose later date as %s or different location as %s, %s", c.mapping.err, c.date, p.Latitude, p.Longitude)
	}
	c.mapping.highTemp, c.mapping.err = c.getTempH()
	if c.mapping.err != nil {
		return fmt.Errorf("%v", c.mapping.err)
	}
	c.mapping.lowTemp, c.mapping.err = c.getTempL()
	if c.mapping.err != nil {
		return fmt.Errorf("%v", c.mapping.err)
	}

	switch p.Fahrenheit {
	case true:
		fmt.Printf("Highest daily temperature was %s Fahrenheit in %s at latitude: %s, longitude %s\n", c.mapping.highTemp, c.date, p.Latitude, p.Longitude)
		fmt.Printf("Lowest daily temperature was %s Fahrenheit in %s at latitude: %s, longitude %s\n", c.mapping.lowTemp, c.date, p.Latitude, p.Longitude)
	default:
		fmt.Printf("Highest daily temperature was %s Celcius in %s at latitude: %s, longitude %s\n", c.mapping.highTemp, c.date, p.Latitude, p.Longitude)
		fmt.Printf("Lowest daily temperature was %s Celcius in %s at latitude: %s, longitude %s\n", c.mapping.lowTemp, c.date, p.Latitude, p.Longitude)
	}
	return nil
}
