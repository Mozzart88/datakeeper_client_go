package datakeeper_client_go

import (
	"net/http"

	api "github.com/Mozzart88/api_go"
)

type DKC struct {
	api.Client
}

var singleton *DKC

// var once sync.Once

func NewDKC(url string, ua string) *DKC {
	if singleton == nil {
		singleton = new(DKC)
		singleton.Url(url)
		singleton.UA(ua)
	}
	return singleton
}

func Get() *DKC {
	return singleton
}

func (dkc *DKC) sendRequest(method string, path string, data api.JsonMap) (*http.Response, error) {
	var c *DKC

	c = Get()
	return c.CallApi(method, path, data)
}

func (dkc *DKC) Get(from string, whereClaus api.JsonMap) (api.Response, error) {
	const method string = "POST"
	var err error
	var response *http.Response
	var res api.Response

	response, err = dkc.sendRequest(method, from, whereClaus)
	if err != nil {
		return res, err
	}
	res = api.NewResponse(response)
	if response.StatusCode == http.StatusOK {
		return res, nil
	}
	if response.StatusCode == http.StatusAccepted {
		return res, nil
	}
	st := response.StatusCode
	msg := res.Body.Get("msg").(string)
	return res, api.NewError(st, msg)
}

func (dkc *DKC) Update(table string, data api.JsonMap) (api.Response, error) {
	const method string = "UPDATE"
	var err error
	var response *http.Response
	var res api.Response

	response, err = dkc.sendRequest(method, table, data)
	if err != nil {
		return res, err
	}
	res = api.NewResponse(response)
	if response.StatusCode == http.StatusOK {
		return res, nil
	}
	if response.StatusCode == http.StatusAccepted {
		return res, nil
	}
	st := response.StatusCode
	msg := res.Body.Get("msg").(string)
	return res, api.NewError(st, msg)
}

func (dkc *DKC) Delete(from string, whereClaus api.JsonMap) (api.Response, error) {
	const method string = "DELETE"
	var err error
	var response *http.Response
	var res api.Response

	response, err = dkc.sendRequest(method, from, whereClaus)
	if err != nil {
		return res, err
	}
	res = api.NewResponse(response)
	if response.StatusCode == http.StatusOK {
		return res, nil
	}
	if response.StatusCode == http.StatusAccepted {
		return res, nil
	}
	st := response.StatusCode
	msg := res.Body.Get("msg").(string)
	return res, api.NewError(st, msg)
}

func (dkc *DKC) Add(to string, data api.JsonMap) (api.Response, error) {
	const method string = "PUT"
	var err error
	var response *http.Response
	var res api.Response

	response, err = dkc.sendRequest(method, to, data)
	if err != nil {
		return res, err
	}
	res = api.NewResponse(response)
	if response.StatusCode == http.StatusOK {
		return res, nil
	}
	if response.StatusCode == http.StatusAccepted {
		return res, nil
	}
	st := response.StatusCode
	msg := res.Body.Get("msg").(string)
	return res, api.NewError(st, msg)
}
