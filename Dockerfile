# Compile stage
FROM golang:1.14-alpine

ENV GOPATH=/
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN go build -o /app main.go

EXPOSE 8080

CMD [ "/app" ]






#
#FROM golang:1.14-alpine AS build-env
#
#ADD . /app
#WORKDIR /app
#
#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download
#
#RUN go build -o /app
#
## Final stage
#FROM debian:buster
#
#EXPOSE 8080
#
#WORKDIR /
#COPY --from=build-env /app /
#
#CMD ["/app"]







#FROM golang:1.14-alpine
#
## Create a directory for the app
#RUN mkdir /app
#
## Copy all files from the current directory to the app directory
#COPY . /app
#
## Set working directory
#WORKDIR /app
#
## Run command as described:
## go build will build an executable file named server in the current directory
#RUN go build -o server .
#
## Run the server executable
#CMD [ "/app/server" ]





#FROM golang:alpine as env
#
## Add support for Delve debugger.
#RUN apk add --no-cache ca-certificates git
#RUN go get github.com/derekparker/delve/cmd/dlv
#
#FROM env as builder
#
#COPY . /app
#RUN go build -o /go/bin/main main.go
#
#FROM alpine as exec
#
#RUN apk add --update bash ca-certificates
#
#WORKDIR /app
#COPY --from=builder /go/bin/main ./
#
#ENTRYPOINT ["/app/main"]






#
#FROM alpine:latest
#
#RUN apk --no-cache add ca-certificates
#WORKDIR /root/
#
#COPY ./.bin/app .
#COPY ./config/ ./config/