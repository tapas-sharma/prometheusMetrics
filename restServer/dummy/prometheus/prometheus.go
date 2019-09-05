package prometheus

import (
	"fmt"
	"net/http"
	"os"

	gokitlog "github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/tapas-sharma/prometheusMetrics/restServer/dummy"
)

var (
	packageName = "prometheus"
	//PingCounter will count the health API call
	PingCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ping",
			Help: "GET on /ping counter",
		})
	//FooCounter will count the health API call
	FooCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "foo",
			Help: "POST on /foo counter",
		})
)

//PromService implements the dummy service
// and also counters for them
type PromService struct {
	hostName string
}

func getHostname() (string, error) {
	name, err := os.Hostname()
	if err != nil {
		logrus.Errorf("[Prometheus] getHostname: %v", err)
	}
	return name, err
}

//Ping for the dummy service
func (ps PromService) Ping(logger gokitlog.Logger) (dummy.PingResponse, error) {
	logger.Log("[Prometheus]", " Ping called")
	resp := dummy.PingResponse{Hostname: ps.hostName, Err: ""}
	PingCounter.Add(1)
	return resp, nil
}

//Foo for the dummy service
func (ps PromService) Foo(req dummy.FooRequest, logger gokitlog.Logger) (dummy.FooResponse, error) {
	msg := fmt.Sprintf("Hello %s, this is BAR from %s!!!", req.Name, ps.hostName)
	resp := dummy.FooResponse{Hostname: msg, Err: ""}
	FooCounter.Add(1)
	return resp, nil
}

//GetCustomRoutes will return prom.http as a endpoint
func (ps PromService) GetCustomRoutes() map[string]http.Handler {
	retval := make(map[string]http.Handler)
	retval["/metrics"] = promhttp.Handler()
	return retval
}

//GetServiceName returns the service registered
func (ps PromService) GetServiceName() string {
	return packageName
}

func init() {
	name, err := getHostname()
	if err != nil {
		panic(err)
	}
	svc := PromService{hostName: name}
	dummy.Register(packageName, svc)
	prometheus.MustRegister(PingCounter)
	prometheus.MustRegister(FooCounter)
}
