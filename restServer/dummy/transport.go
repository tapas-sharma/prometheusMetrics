package dummy

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	gokitlog "github.com/go-kit/kit/log"
	gokithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

//Route can be used to define custom Routes
// in the dummy service
type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc *gokithttp.Server
}

//Routes that need to be added
type routes []route

//Logger for API's
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		//inner is like func in Python
		inner.ServeHTTP(w, r)
		//how do we log in syslog?
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// MakeHandler creates and endpoint for create role
func MakeHandler(bs Service, logger gokitlog.Logger, r *mux.Router) http.Handler {

	opts := []gokithttp.ServerOption{
		gokithttp.ServerErrorLogger(logger),
		gokithttp.ServerErrorEncoder(dummyErrorEncoder),
	}

	pingHandler := gokithttp.NewServer(
		makePingEndpoint(bs, logger),
		decodePingRequest,
		encodeResponse,
		opts...,
	)

	fooHandler := gokithttp.NewServer(
		makeFooEndpoint(bs, logger),
		decodeFooRequest,
		encodeResponse,
		opts...,
	)

	var routes = routes{
		route{
			"PingRequest",
			"GET",
			"/ping",
			pingHandler,
		},
		route{
			"FooRequest",
			"POST",
			"/foo",
			fooHandler,
		},
	}

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return r
}

func dummyErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(PingResponse{Hostname: "", Err: err.Error()})
}

func decodePingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeFooRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request FooRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
