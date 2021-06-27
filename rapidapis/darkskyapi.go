package rapidapis

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"weather-api/defaults"
	"weather-api/utils"
)

func DarkSkyC2(p Params) (*DarkSkyClient, error) {
	date, err := (&p).validateParams()
	if err != nil {
		log.Fatalf("❌ Parameter validation failed: %v", err)
	}
	url := utils.DarkSkyBuildBaseURL(p.Latitude, p.Longitude, date)
	request, err := http.NewRequest(defaults.GET, url, nil)
	if err != nil {
		return nil, fmt.Errorf("eror when creating http GET request, error:%v", err)
	}
	request.Header.Add(defaults.RapidApiHeaderKey, p.Apikey)
	request.Header.Add(defaults.RapidApiHeaderHost, defaults.DarkSkyApi)

	var httpsClient = &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := httpsClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error when getting http GET response, error:%v", err)
	}
	defer response.Body.Close()
	data, err := utils.JsonDecoder(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error when decoding http.Request.Body to json")
	}
	return &DarkSkyClient{
		data: data,
		err:  err,
		date: date,
	}, nil
}

func DarkSkyreturns2(p Params) error {
	c, err := DarkSkyC2(p)
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
		return fmt.Errorf("error occured with highTemp,error:%v", c.mapping.err)
	}
	c.mapping.lowTemp, c.mapping.err = c.getTempL()
	if c.mapping.err != nil {
		return fmt.Errorf("error occured with LowTemp,error:%v", c.mapping.err)
	}
	fmt.Printf("Highest daily temperature was %s Celcius in %s in %s, %s\n", c.mapping.highTemp, c.date, p.City, c.countryCode)
	fmt.Printf("Lowest daily temperature was %s Celcius in %s in %s, %s\n", c.mapping.lowTemp, c.date, p.City, c.countryCode)
	return nil
}
