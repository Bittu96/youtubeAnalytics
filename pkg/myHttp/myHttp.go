package myHttp

import (
	"io"
	"log"
	"net/http"
	"reflect"
	"time"
	"youtubeAnalytics/pkg/logger"
)

type request struct {
	method string
	url    string
	params interface{}
}

const maxRetryCount = 10
const retryInterval = time.Second

func NewRequest(method, url string, params interface{}) request {
	return request{
		method: method,
		url:    url,
		params: params,
	}
}

func (r request) Do() ([]byte, error) {
	// create http request from params
	req, err := r.createHTTPRequest()
	if err != nil {
		return nil, err
	}

	// do http call with retries
	res, err := doWithRetries(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r request) createHTTPRequest() (req *http.Request, err error) {
	req, err = http.NewRequest(r.method, r.url, nil)
	if err != nil {
		return
	}
	queryParams := req.URL.Query()

	// add query params to request
	values := reflect.ValueOf(r.params)
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {
		key := types.Field(i).Tag.Get("json")
		value := values.Field(i).String()
		queryParams.Add(key, value)
	}

	req.URL.RawQuery = queryParams.Encode()
	// fmt.Println(req.URL.String())
	return
}

func doWithRetries(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	res, err := client.Do(req)
	for i := 0; (i < maxRetryCount) && shouldRetryConnection(res, err); i++ {
		log.Println("retrying youtube api call...", retryInterval.String())
		time.Sleep(retryInterval)
		res, err = client.Do(req)
		if err != nil {
			logger.Error(err.Error())
		}
	}

	return res, err
}

// check retry condition
func shouldRetryConnection(res *http.Response, err error) bool {
	if err != nil {
		return true
	}
	if res.StatusCode == http.StatusBadGateway || res.StatusCode == http.StatusServiceUnavailable || res.StatusCode == http.StatusGatewayTimeout {
		return true
	}
	return false
}
