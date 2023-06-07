FROM golang:1.20.5-alpine3.18

RUN adduser -D scotch

RUN apk update

WORKDIR /go/app

COPY . .

RUN chown -R scotch:scotch .

USER scotch

RUN go install github.com/cosmtrek/air@latest && go get .

EXPOSE 80

ENTRYPOINT ["air"]
