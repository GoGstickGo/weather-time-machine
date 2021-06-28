package clients

import (
	"fmt"
	"net/http"
	"time"
	"weather-api/defaults"
)

func CreateClient(apikey, apimethod, apiurl, baseurl string) (*http.Response, error) {
	request, err := http.NewRequest(apimethod, baseurl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http GET request, error:%v", err)
	}
	request.Header.Add(defaults.RapidApiHeaderKey, apikey)
	request.Header.Add(defaults.RapidApiHeaderHost, apiurl)

	var httpsClient = &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := httpsClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to create http GET response, error:%v", err)
	}
	return response, nil
}
