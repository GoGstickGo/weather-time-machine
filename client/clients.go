package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"weather-api/defaults"
)

type Responses struct {
	Method     string
	Url        string
	Apiaddress string
	Request    *http.Request
	Response   *http.Response
	Data       map[string]interface{}
	Error      error
}

func jsonDecoder(r io.Reader) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	return data, json.NewDecoder(r).Decode(&data)
}

func Client(r *Responses) (map[string]interface{}, error) {
	r.Request, r.Error = http.NewRequest(r.Method, r.Url, nil)
	if r.Error != nil {
		return r.Data, fmt.Errorf("eror when creating http GET request, error:%v", r.Error)
	}
	r.Request.Header.Add(defaults.RapidApiHeaderKey, defaults.RapidApiKey)
	r.Request.Header.Add(defaults.RapidApiHeaderHost, r.Apiaddress)
	// add timeout
	var httpsClient = &http.Client{
		Timeout: time.Second * 10,
	}
	r.Response, r.Error = httpsClient.Do(r.Request)
	if r.Error != nil {
		return r.Data, fmt.Errorf("error when getting http GET response, error:%v", r.Error)
	}
	r.Data, r.Error = jsonDecoder(r.Response.Body)
	if r.Error != nil {
		return r.Data, fmt.Errorf("error when decoding http.Request.Body to json")
	}
	return r.Data, nil
}
