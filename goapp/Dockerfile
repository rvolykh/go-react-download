FROM golang:1.7-alpine3.6 AS build

COPY . /go/src/github.com/rvolykh/react-go-download/goapp

RUN go install github.com/rvolykh/react-go-download/goapp

FROM alpine:3.6
COPY --from=build /go/bin/goapp /opt/goapp
RUN chmod +x /opt/goapp
EXPOSE 80
ENTRYPOINT /opt/goapp
