FROM golang:1.22 as base

WORKDIR /app

COPY mqtt_distributer/go.mod mqtt_distributer/go.sum ./

RUN go mod download

COPY mqtt_distributer/main.go ./

RUN go build -o main .

CMD ["./main"]
