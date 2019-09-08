# Prometheus Metrics Server

The code in this repo tries to implement a dummy server that has the following REST API's

|API|Method|URL|Curl|
|-|-|-|-|
|ping|GET|http://<host>:<port>/ping|curl -XGET http://localhost:8080/ping|
|foo|POST|http://<host>:<port>/ping|curl -XPOST -d '{"name":"FOO"}' http://localhost:8080/foo|

You can also implement a set of custom REST endpoint that can be returned using the `GetCustomRoutes()`.
The current `dummy` rest server is implemented by the `prometheus` package with the above 2 REST endpoint and also the `/metrics` endpoint needed by the Prometheus service to scrape counter.
The prometheus `dummy` server exposes the following counters
* ping 
* foo 

Which can be used to show the total count of requests that were received for these endpoints.
The code uses gorilla-mux and we could have used the sdk-kit's kitpromtheus to do the same, thus reducing code, but I was trying to see how we can write better REST API's in goLang.

## How to compile
#### To compile
* Setup your GOPATH and GOBIN correctly 
```
export GOPATH=~/go
export GOBIN=$GOPATH/bin/
export PATH=$PATH:$GOBIN
```
* Then run make
```
make
```
This will install `glide` check the dependencies and create binaries for the rest server under the following path
```
$GOPATH/github.com/tapas-sharma/prometheusMetrics/restServer/bin
```
Choose the correct platform you are running in and cd into it, i.e
* linux
* osx
* windows

and then run `restServer(.exe)`

## Docker image
To run the code without compiling
```
docker run --rm -d -p 8080:8080 tapassharma/prometheus-metrics:v1 
```
