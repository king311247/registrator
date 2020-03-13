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
	return &HttpcollectorAdapter{client: http.DefaultClient}
}

type HttpcollectorAdapter struct {
	client *http.Client
}

func (h HttpcollectorAdapter) Ping() error {
	log.Println("httpcollector ping ")
	return nil
}

func (h HttpcollectorAdapter) Register(service *bridge.Service) error {
	data, _ := json.Marshal(service)
	jsonStr := string(data)
	log.Println("Register : " + jsonStr)
	return nil
}

func (h HttpcollectorAdapter) Deregister(service *bridge.Service) error {
	data, _ := json.Marshal(service)
	jsonStr := string(data)
	log.Println("Deregister : " + jsonStr)
	return nil
}

func (h HttpcollectorAdapter) Refresh(service *bridge.Service) error {
	return nil
}

func (h HttpcollectorAdapter) Services() ([]*bridge.Service, error) {
	return nil, nil
}
