FROM golang:1.19

WORKDIR /var/www/api

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum main.go .air.toml .

EXPOSE 3000

CMD ["air", "-c", ".air.toml"]
