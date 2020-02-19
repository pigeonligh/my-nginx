FROM golang:alpine AS building

LABEL maintainer="cx24321@hotmail.com"

ENV GOPATH=/go

RUN apk add --update -t build-deps curl go git libc-dev gcc libgcc

RUN go get -u -v github.com/gin-gonic/gin

WORKDIR /go/src/github.com/pigeonligh/my-nginx

COPY . .

RUN go build -o main

FROM nginx:alpine

WORKDIR /go/src/github.com/pigeonligh/my-nginx

COPY --from=building /go/src/github.com/pigeonligh/my-nginx .

EXPOSE 8000

CMD ["./main"]
