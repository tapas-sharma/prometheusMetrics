package dummy

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	gokitlog "github.com/go-kit/kit/log"
)

func makePingEndpoint(svc Service, logger gokitlog.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.Log("Service", svc.GetServiceName(), "Endpoint", "Ping")
		return svc.Ping(logger)
	}
}

func makeFooEndpoint(svc Service, logger gokitlog.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FooRequest)
		logger.Log("Service", svc.GetServiceName(), "Endpoint", "Foo")
		return svc.Foo(req, logger)
	}
}
