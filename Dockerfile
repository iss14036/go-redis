FROM golang:latest

RUN apt-get update
RUN apt-get upgrade -y

ENV GO111MODULE=on

COPY . /go-redis/

WORKDIR /go-redis

RUN go get -u github.com/go-redis/redis/v7


CMD ["go" ,"run", "main.go"]
