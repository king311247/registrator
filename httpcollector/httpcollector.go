package httpcollector

import (
	"encoding/json"
	"github.com/gliderlabs/registrator/bridge"
	"log"
	"net/http"
	//"fmt"
	//"log"
	"net/url"
)

const DefaultInterval = "10s"

func init() {
	f := new(Factory)
	bridge.Register(f, "httpcollector")
}

type Factory struct{}

func (f *Factory) New(uri *url.URL) bridge.RegistryAdapter {
	return &HttpCollectAdapter{client: http.DefaultClient}
}

type HttpCollectAdapter struct {
	client *http.Client
}

func (h HttpCollectAdapter) Ping() error {
	log.Println("httpcollector ping ")
	return nil
}

func (h HttpCollectAdapter) Register(service *bridge.Service) error {
	data, _ := json.Marshal(service)
	jsonStr := string(data)
	log.Println("Register : " + jsonStr)
	return nil
}

func (h HttpCollectAdapter) Deregister(service *bridge.Service) error {
	data, _ := json.Marshal(service)
	jsonStr := string(data)
	log.Println("Deregister : " + jsonStr)
	return nil
}

func (h HttpCollectAdapter) Refresh(service *bridge.Service) error {
	return nil
}

func (h HttpCollectAdapter) Services() ([]*bridge.Service, error) {
	return nil, nil
}
