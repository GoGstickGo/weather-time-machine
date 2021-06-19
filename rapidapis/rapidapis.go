package rapidapis

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	"weather-api/defaults"
	"weather-api/utils"
)

type Params struct {
	Day    string
	Month  string
	Year   string
	Apikey string
	City   string
	Writer io.Writer
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
		return c.mapping.latitude, c.mapping.longitude, fmt.Errorf("failed to get city cordinates")
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
		return "", fmt.Errorf("invalid date form %s,%v", date, err)
	}
	return date, nil
}

func geoDBClient(p Params) (*GeoDBClient, error) {
	_, err := (&p).validateParams()
	if err != nil {
		log.Fatalf("❌ Parameter validation failed: %v", err)
	}

	url := utils.GeoDBBuildBaseURL(p.City)
	request, err := http.NewRequest(defaults.GET, url, nil)
	if err != nil {
		return nil, fmt.Errorf("eror when creating http GET request, error:%v", err)
	}
	request.Header.Add(defaults.RapidApiHeaderKey, p.Apikey)
	request.Header.Add(defaults.RapidApiHeaderHost, defaults.GeoDBApi)
	// add timeout
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

func GeoDBreturns(p Params) (string, string, string, error) {
	c, err := geoDBClient(p)
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

func DarkSkyC(p Params) (*DarkSkyClient, error) {
	countryCode, latitude, longitude, err := GeoDBreturns(p)
	if err != nil {
		log.Fatalf("❌ Failed to get values from GeoDB Api %v", err)
	}
	date, _ := (&p).validateParams()

	url := utils.DarkSkyBuildBaseURL(latitude, longitude, date)
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
		data:        data,
		err:         err,
		date:        date,
		countryCode: countryCode,
	}, nil
}

func DarkSkyreturns(p Params) error {
	c, err := DarkSkyC(p)
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
