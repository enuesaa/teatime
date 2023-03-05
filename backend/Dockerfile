FROM golang:1.19 AS base

WORKDIR /app

COPY . .
RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]


FROM base AS dev

RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

RUN apt update \
    && apt install -y nodejs npm
RUN npm install -g @bufbuild/protoc-gen-es 
RUN npm install -g @bufbuild/protoc-gen-connect-web

CMD ["air", "-c", ".air.toml"]
