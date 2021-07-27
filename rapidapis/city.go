package rapidapis

import (
	"fmt"
	"log"
	"weather-api/clients"
	"weather-api/defaults"
	"weather-api/utils"
)

func (p *Params) validateParams() (string, error) {
	err := utils.ValidateParamsCity(p.City)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	err = utils.ValidateParamsApikey(p.Apikey)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	date, err := utils.BuildDate(p.Year, p.Month, p.Day)
	if err != nil {
		return "", fmt.Errorf("invalid date form %s, %v", date, err)
	}
	return date, nil
}

func gdClient(p Params) (*GeoDBClient, error) {
	_, err := (&p).validateParams()
	if err != nil {
		log.Fatalf("❌ Parameter validation failed: %v", err)
	}

	url := utils.GeoDBBuildBaseURL(p.City)

	response, err := clients.CreateClient(p.Apikey, defaults.GET, defaults.GeoDBApi, url)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	data, err := utils.JsonDecoder(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error when decoding http.Request.Body to json")
	}

	err = utils.ValidateRapidApiKey(data)
	if err != nil {
		log.Fatalf("❌ RapidApi error: %v", err)
	}
	return &GeoDBClient{
		data:   data,
		err:    err,
		params: p,
	}, nil
}

func gdReturns(p Params) (string, string, string, error) {
	c, err := gdClient(p)
	if err != nil {
		log.Fatalf("❌ Error occured establishing client for GeoDB %v", err)
	}
	c.mapping.recieved = c.data
	err = c.validateCity()
	if err != nil {
		log.Fatalf("❌ GeoDB cities API error: %v", err)
	}
	c.mapping.tempField, c.mapping.err = c.convertMap()
	if c.mapping.err != nil {
		return "", "", "", fmt.Errorf("convert maps: %v", c.mapping.err)
	}
	c.mapping.latitude, c.mapping.longitude, c.mapping.err = c.getCityLocation()
	if c.mapping.err != nil {
		return "", "", "", fmt.Errorf("when getting cordinates for the %s", c.params.City)
	}
	c.mapping.countryCode, c.mapping.err = c.getCountryCode()
	if c.mapping.err != nil {
		return "", "", "", fmt.Errorf("when getting countryCode for the %s, error:%v", c.params.City, c.mapping.err)
	}
	return c.mapping.countryCode, c.mapping.latitude, c.mapping.longitude, nil
}

func dsClient(p Params) (*DarkSkyClient, error) {
	var url string

	countryCode, latitude, longitude, err := gdReturns(p)
	if err != nil {
		log.Fatalf("❌ Failed to get values from GeoDB Api %v", err)
	}

	date, _ := (&p).validateParams()

	switch p.Fahrenheit {
	case true:
		url = utils.DarkSkyBuildBaseURLFahrenheit(latitude, longitude, date)
	default:
		url = utils.DarkSkyBuildBaseURLCelcius(latitude, longitude, date)
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
		data:        data,
		err:         err,
		date:        date,
		countryCode: countryCode,
	}, nil
}

func DsReturns(p Params) error {
	c, err := dsClient(p)
	if err != nil {
		log.Fatalf("❌ Error occured establishing client for DarkSky Api %v", err)
	}
	c.mapping.recieved = c.data
	c.mapping.tempField, c.mapping.err = c.convertMap()
	if c.mapping.err != nil {
		log.Fatalf("❌ DarkSkyApi error: %v, please choose later date than %s", c.mapping.err, c.date)
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
		fmt.Printf("Highest daily temperature was %s Fahrenheit in %s in %s, %s\n", c.mapping.highTemp, c.date, p.City, c.countryCode)
		fmt.Printf("Lowest daily temperature was %s Fahrenheit in %s in %s, %s\n", c.mapping.lowTemp, c.date, p.City, c.countryCode)
	default:
		fmt.Printf("Highest daily temperature was %s Celcius in %s in %s, %s\n", c.mapping.highTemp, c.date, p.City, c.countryCode)
		fmt.Printf("Lowest daily temperature was %s Celcius in %s in %s, %s\n", c.mapping.lowTemp, c.date, p.City, c.countryCode)
	}
	return nil
}
