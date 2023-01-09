FROM golang:1.19

WORKDIR /var/www/api

COPY go.mod go.sum main.go .air.toml .

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
