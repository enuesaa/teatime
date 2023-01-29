FROM golang:1.19

WORKDIR /app

COPY . .
RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
