package main

import (
	"fmt"
	"log"
	c "weather-api/client"
	d "weather-api/defaults"
	u "weather-api/utils"
)

func main() {

	/*year := "2018"
	month := "07"
	day := "27"
	city := "Dublin"*/

	/*dataTime, err := u.BuildTimeStr(year, month, day)
	if err != nil {
		log.Fatalf("error occured with date in the API request: %v", err)
	}

	alti, longi, err := u.CityLocation(city)
	if err != nil {
		log.Fatalf("error occured with city corditnates %v", err)
	}*/
	client := c.Responses{
		//Url:        d.DarkSkyApiUrl + alti + "," + longi + "," + dataTime + d.DarkSkyApiSort,
		Url:        "https://dark-sky.p.rapidapi.com/46.520277777,17.020277777,2018-07-27T12:00:00?units=ca&exclude=currently%2Chourly%2Calerts%2Cflags&lang=en",
		Method:     "GET",
		Apiaddress: d.DarkSkyApi,
	}

	client.Request, client.Error = c.HttpGetRequest(&client)
	if client.Error != nil {
		log.Fatalf("error occured with httpGetRequest: %v", client.Error)
	}

	/*req, err := c.("GET", url, "dark-sky.p.rapidapi.com", nil)
	if err != nil {
		log.Fatalf("error occured with httpGetRequest: %v", err)
	}*/

	client.Response, client.Error = c.HttpGetResponse(&client)
	if client.Error != nil {
		log.Fatalf("error occured with httpGetResponse: %v", client.Error)
	}

	defer client.Response.Body.Close()
	body, err := u.JsonDecoder(client.Response.Body)
	if err != nil {
		log.Fatalf("cannot decode Json , error: %v", err)
	}

	convMap, err := u.ConvertMapTemp(body)
	if err != nil {
		log.Fatalf("error occured with response body: %v", err)
	}

	/*printDate, err := u.PrintDate(dataTime)
	if err != nil {
		fmt.Printf("Something wrong with date in the output.")
	}*/

	highTemp, err := u.GetHighTemp(convMap)
	if err != nil {
		log.Fatalf("error occured during getting the value for tempareture: %v", err)
	}

	fmt.Printf("%s highest temperature of the day was %s Celsius at %s\n", "city", highTemp, "something")
	lowTemp, err := u.GetLowTemp(convMap)
	fmt.Println(lowTemp)
	if lowTemp == "" || highTemp == "" {
		log.Fatalf("there is no data for date : %s, please choose another date", "something")
	}

	if err != nil {
		log.Fatalf("error occured during getting the value for tempareture: %v", err)

	}

	fmt.Printf("%s lowest temperature of the day was %s Celsius at %s\n", "city", lowTemp, "something")
}
