FROM golang:latest

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

COPY go.mod /go/src/app
# COPY go.sum /go/src/app
# COPY . /go/src/app

RUN go mod download

RUN apt-get update && apt-get install git

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
