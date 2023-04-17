FROM golang:1.20.3-alpine3.17

WORKDIR app

COPY . ./

RUN apk update && go install github.com/cosmtrek/air@latest && go get .

ENTRYPOINT ["air"]
