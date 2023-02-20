FROM golang:1.19

WORKDIR /app

COPY . .
RUN go install github.com/cosmtrek/air@latest

# tmp
RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

CMD ["air", "-c", ".air.toml"]
