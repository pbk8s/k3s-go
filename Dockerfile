FROM golang:1.11-alpine as builder
#FROM golang:1.14-alpine as builder
#FROM arm64v8/golang:1.14-alpine as builder
WORKDIR /usr/build
#RUN apt-get install build-essential -qy
ADD main.go .
RUN go build -o app .

FROM alpine:latest
#FROM arm64v8/alpine:latest

WORKDIR /usr/src

COPY --from=builder /usr/build/app .
EXPOSE 8080

CMD ["/usr/src/app"]
