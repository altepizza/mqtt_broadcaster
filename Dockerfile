FROM golang:1.22 as base

LABEL org.opencontainers.image.source https://github.com/altepizza/mqtt_broadcaster

WORKDIR /app

COPY mqtt_distributer/go.mod mqtt_distributer/go.sum ./

RUN go mod download

COPY mqtt_distributer/main.go ./

RUN go build -o main .

CMD ["./main"]
