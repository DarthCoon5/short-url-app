FROM golang:1.14.6-alpine3.12 as builder

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o short-url-app ./main.go

CMD ["./short-url-app"]


#
#FROM alpine:latest
#
#RUN apk --no-cache add ca-certificates
#WORKDIR /root/
#
#COPY ./.bin/app .
#COPY ./config/ ./config/