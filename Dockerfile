FROM golang:1.10.3-alpine
RUN apk update && apk add --no-cache ca-certificates
RUN echo 1
