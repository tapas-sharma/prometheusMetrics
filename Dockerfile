FROM golang:1.12-alpine as builder

RUN mkdir -p $GOPATH/src/github.com/tapas-sharma/prometheusMetrics && apk add --update make git curl 
RUN mkdir -p $GOPATH/bin
ENV GOBIN=$GOPATH/bin
RUN curl https://glide.sh/get | sh
ADD . $GOPATH/src/github.com/tapas-sharma/prometheusMetrics
RUN cd $GOPATH/src/github.com/tapas-sharma/prometheusMetrics && make

FROM alpine:latest
COPY --from=builder /go/src/github.com/tapas-sharma/prometheusMetrics/restServer/bin/linux/restServer /usr/local/bin
EXPOSE 8080
ENTRYPOINT [ "restServer"]