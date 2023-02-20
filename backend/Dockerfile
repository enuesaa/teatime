FROM golang:1.19

WORKDIR /app

COPY . .
RUN go install github.com/cosmtrek/air@latest

RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

CMD ["air", "-c", ".air.toml"]
