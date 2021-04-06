FROM golang:1.16.3-alpine3.13 as builder

ENV GOFLAGS=-mod=vendor

COPY . /go/src/github.com/quintsys/ga-exporter/
WORKDIR /go/src/github.com/quintsys/ga-exporter

RUN CGO_ENABLED=0 go build -a -v

FROM alpine:3.13
COPY --from=builder /go/src/github.com/quintsys/ga-exporter/ga-exporter /opt/

EXPOSE 8080
CMD ["/opt/ga-exporter"]
