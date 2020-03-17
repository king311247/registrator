package httpcollector

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/king311247/registrator/bridge"
	"io/ioutil"
	"log"
	"net/http"
	//"fmt"
	"net/url"
)

const DefaultInterval = "10s"

func init() {
	f := new(Factory)
	bridge.Register(f, "httpcollector")
}

type Factory struct{}

func (f *Factory) New(uri *url.URL) bridge.RegistryAdapter {
	return &HttpcollectorAdapter{client: http.DefaultClient, baseUrl: "http://" + uri.Host}
}

type HttpcollectorAdapter struct {
	client  *http.Client
	baseUrl string
}

func (h HttpcollectorAdapter) Ping() error {
	var url = h.baseUrl + "/api/serviceregister/ping"
	response, err := h.client.Get(url)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New("response status " + response.Status)
	}
	return nil
}

func (h HttpcollectorAdapter) Register(service *bridge.Service) error {
	postData, err := json.Marshal(service)
	if err != nil {
		return err
	}

	log.Println("Register : " + string(postData))

	var url = h.baseUrl + "/api/serviceregister/register"
	response, err := h.client.Post(url, "application/json", bytes.NewReader(postData))
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New("response status " + response.Status)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	apiResponse := new(ReregisterResponse)
	err = json.Unmarshal(body, apiResponse)
	if err != nil {
		return errors.New("服务端返回：" + string(body))
	}

	if apiResponse.Code != 0 {
		return errors.New("服务端返回：" + string(body))
	}

	return nil
}

func (h HttpcollectorAdapter) Deregister(service *bridge.Service) error {
	postData, err := json.Marshal(service)

	if err != nil {
		return err
	}

	log.Println("Deregister : " + string(postData))

	var url = h.baseUrl + "/api/serviceregister/deregister"
	response, err := h.client.Post(url, "application/json", bytes.NewReader(postData))
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New("response status " + response.Status)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	apiResponse := new(ReregisterResponse)
	err = json.Unmarshal(body, apiResponse)
	if err != nil {
		return errors.New("服务端返回：" + string(body))
	}

	if apiResponse.Code != 0 {
		return errors.New("服务端返回：" + string(body))
	}

	return nil
}

func (h HttpcollectorAdapter) Refresh(service *bridge.Service) error {
	return nil
}

// 注册、注销请求响应
type ReregisterResponse struct {
	Code    int
	Message string
}

// 服务端返回的数据结构
type ApiService struct {
	ID      string
	Service string
	Port    int
	Tags    []string
	Address string
}

// 服务列表请求响应
type ApiServicesResponse struct {
	Code    int
	Message string
	Data    []ApiService
}

func (h HttpcollectorAdapter) Services() ([]*bridge.Service, error) {
	var url = h.baseUrl + "/api/serviceregister/servicelist"
	response, err := h.client.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("response status " + response.Status)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := new(ApiServicesResponse)
	err = json.Unmarshal(body, apiResponse)
	if err != nil {
		return nil, errors.New("服务端返回：" + string(body))
	}

	if apiResponse.Code != 0 {
		return nil, errors.New("服务端返回：" + string(body))
	}

	out := make([]*bridge.Service, len(apiResponse.Data))
	i := 0

	for _, v := range apiResponse.Data {
		s := &bridge.Service{
			ID:   v.ID,
			Name: v.Service,
			Port: v.Port,
			Tags: v.Tags,
			IP:   v.Address,
		}
		out[i] = s
		i++
	}
	return out, nil
}
