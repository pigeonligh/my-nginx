FROM golang:1.13.9-alpine3.11 AS building

LABEL maintainer="cx24321@hotmail.com"

ENV GOPATH=/go

RUN apk add --update -t build-deps curl go git libc-dev gcc libgcc
RUN go get -u -v github.com/gin-gonic/gin

WORKDIR /go/src/github.com/pigeonligh/my-nginx
COPY . .
RUN go build -o nginx-manager

FROM nginx:1.13.9-alpine3.11
WORKDIR /opt/app
COPY --from=building /go/src/github.com/pigeonligh/my-nginx .

CMD [ "sh", "-c", "./nginx-manager -token $MANAGE_TOKEN" ]

EXPOSE 8080
EXPOSE 80
EXPOSE 443
