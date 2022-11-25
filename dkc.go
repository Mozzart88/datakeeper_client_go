package datakeeper_client_go

import (
	"strings"
	"sync"

	api "github.com/Mozzart88/api_go"
)

type DKC api.Server

var singleton *DKC
var once sync.Once

func NewClient(url string, secrete string, ua string) *DKC {
	once.Do(func() {
		singleton = new(DKC)
		singleton.Secret = secrete
		singleton.Url = url
		singleton.UserAgent = ua
	})
	return singleton
}

func Get() *DKC {
	if singleton == nil {
		singleton = NewClient("", "", "")
	}
	return singleton
}

func (dkc *DKC) newServer(path string) *api.Server {
	var s *api.Server

	s = new(api.Server)
	s.Secret = dkc.Secret
	s.Url = strings.Join([]string{dkc.Url, path}, "/")
	s.UserAgent = dkc.UserAgent
	return s
}

func (dkc *DKC) sendRequest(method string, path string, data api.JsonMap) (api.ApiResponse, error) {
	var err error
	var res api.ApiResponse
	var s *api.Server

	s = dkc.newServer(path)
	res, err = s.CallApi(method, data)
	return res, err
}

func (dkc *DKC) Get(from string, whereClaus api.JsonMap) (api.ApiResponse, error) {
	var err error
	var res api.ApiResponse

	res, err = dkc.sendRequest("POST", from, whereClaus)
	return res, err
}

func (dkc *DKC) Update(table string, data api.JsonMap) (api.ApiResponse, error) {
	var err error
	var res api.ApiResponse

	res, err = dkc.sendRequest("UPDATE", table, data)
	return res, err
}

func (dkc *DKC) Delete(from string, whereClaus api.JsonMap) (api.ApiResponse, error) {
	var err error
	var res api.ApiResponse

	res, err = dkc.sendRequest("DELETE", from, whereClaus)
	return res, err
}

func (dkc *DKC) Add(to string, data api.JsonMap) (api.ApiResponse, error) {
	var err error
	var res api.ApiResponse

	res, err = dkc.sendRequest("PUT", to, data)
	return res, err
}
