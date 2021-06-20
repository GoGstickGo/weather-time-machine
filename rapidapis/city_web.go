package rapidapis

import (
	"fmt"
	"weather-api/clients"
	"weather-api/defaults"
	"weather-api/utils"
)

func (p *Params) validateParamsWeb() (string, error) {
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

func gdClientWeb(p Params) (*GeoDBClient, error) {
	_, err := (&p).validateParamsWeb()
	if err != nil {
		return nil, fmt.Errorf("%v", err)
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
		return nil, fmt.Errorf("%v", err)
	}
	return &GeoDBClient{
		data:   data,
		err:    err,
		params: p,
	}, nil
}

func gdReturnsWeb(p Params) (string, string, string, error) {
	c, err := gdClientWeb(p)
	if err != nil {
		return "", "", "", fmt.Errorf("%v", err)
	}
	c.mapping.recieved = c.data
	err = c.validateCity()
	if err != nil {
		return "", "", "", fmt.Errorf("%v", err)
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

func dsClientWeb(p Params) (*DarkSkyClient, error) {
	countryCode, latitude, longitude, err := gdReturnsWeb(p)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	date, _ := (&p).validateParams()

	url := utils.DarkSkyBuildBaseURL(latitude, longitude, date)

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

func DsReturnsWeb(p Params) (tempH, tempL, date, name, cc string, e error) {
	c, err := dsClientWeb(p)
	if err != nil {
		return "", "", "", "", "", fmt.Errorf("%v", err)
	}
	c.mapping.recieved = c.data
	c.mapping.tempField, c.mapping.err = c.convertMap()
	if c.mapping.err != nil {
		return "", "", "", "", "", fmt.Errorf("%v, please choose later date than %s", c.mapping.err, c.date)
	}
	c.mapping.highTemp, c.mapping.err = c.getTempH()
	if c.mapping.err != nil {
		return "", "", "", "", "", fmt.Errorf("%v", c.mapping.err)
	}
	c.mapping.lowTemp, c.mapping.err = c.getTempL()
	if c.mapping.err != nil {
		return "", "", "", "", "", fmt.Errorf("%v", c.mapping.err)
	}
	return c.mapping.highTemp, c.mapping.lowTemp, c.date, p.City, c.countryCode, nil
}
