FROM golang:latest

WORKDIR /var/www/api

# see https://github.com/oxequa/realize/issues/269
RUN GO111MODULE=off go get github.com/oxequa/realize

CMD ["realize", "start", "--run"]