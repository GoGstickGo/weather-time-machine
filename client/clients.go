package client

import (
	"fmt"
	"net/http"
	"time"
)

func HttpGetRequest(m string, url string, api string, e error) (*http.Request, error) {
	req, err := http.NewRequest(m, url, nil)
	if err != nil {
		return req, fmt.Errorf("eror when creating http GET request, error:%v", err)
	}
	req.Header.Add("x-rapidapi-key", "110f37f7cemshb593a342050b246p15da04jsn37a3042ae0e4")
	req.Header.Add("x-rapidapi-host", api)
	return req, nil
}

func HttpGetResponse(r *http.Request) (*http.Response, error) {
	//Add timeout
	var httpsClient = &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := httpsClient.Do(r)
	if err != nil {
		return res, fmt.Errorf("error when getting http GET response, error:%v", err)
	}
	return res, nil
}
