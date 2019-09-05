package dummy

import (
	"fmt"
	"net/http"

	gokitlog "github.com/go-kit/kit/log"
)

//Service defines the bare minimum functions
// that should be implemented by any struct that wants
// to use this interface
type Service interface {
	Ping(logger gokitlog.Logger) (PingResponse, error)
	Foo(req FooRequest, logger gokitlog.Logger) (FooResponse, error)
	GetCustomRoutes() map[string]http.Handler
	GetServiceName() string
}

// PingResponse for /ping endpoint
type PingResponse struct {
	Hostname string `json:"hostname,omitempty"` //contains the hostname of the machine we are running on
	Err      string `json:"err,omitempty"`      //Actual Error message
}

// FooRequest for the /foo endpoint
type FooRequest struct {
	Name string `json:"name,omitempty"` //name if any
}

// FooResponse for the /foo endpoint
type FooResponse struct {
	Hostname string `json:"hostname,omitempty"` //contains message that we need to send to bar :-)
	Err      string `json:"err,omitempty"`      //Actual Error message
}

var dummyServices = make(map[string]Service)

//Register the Dummy Service with its Service implementation
func Register(name string, svc Service) error {
	dummyServices[name] = svc
	return nil
}

//Get the implementation for the Service
func Get(name string) (Service, error) {
	svc, ok := dummyServices[name]
	if !ok {
		return nil, fmt.Errorf("Get: Service %s not registered", name)
	}
	return svc, nil
}
