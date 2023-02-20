FROM golang:1.19

WORKDIR /app

COPY . .
RUN go install github.com/cosmtrek/air@latest
# tmp
RUN go install github.com/swaggo/swag/cmd/swag@latest

CMD ["air", "-c", ".air.toml"]
