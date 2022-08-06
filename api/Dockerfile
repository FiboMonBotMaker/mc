FROM golang:1.18-rc-alpine3.15

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git



COPY ./app ${ROOT}

RUN go mod tidy

RUN go get github.com/go-sql-driver/mysql
RUN go get xorm.io/xorm
RUN go get github.com/joho/godotenv
