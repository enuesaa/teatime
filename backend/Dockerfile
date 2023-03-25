FROM golang:1.20 AS production

WORKDIR /app
COPY . .
RUN go build -o main
ENTRYPOINT ["/app/main"]


FROM golang:1.20 AS develop

WORKDIR /workspace
RUN apt update
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/magefile/mage@latest
RUN go install github.com/mgechev/revive@latest

RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

RUN apt install -y nodejs npm
RUN npm install -g @bufbuild/protoc-gen-es 
RUN npm install -g @bufbuild/protoc-gen-connect-web

WORKDIR /workspace/backend

CMD ["air", "-c", ".air.toml"]
