FROM golang:1.16 as base

FROM base as dev
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
WORKDIR /opt/app/api
CMD ["air"]

FROM base as built
WORKDIR /go/app/api
COPY . .
ENV CGO_ENABLED=0
RUN go get -d -v ./...
RUN go build -o /tmp/main ./*.go

FROM alpine
COPY --from=built /tmp/main /usr/bin/bot
CMD ["bot"]