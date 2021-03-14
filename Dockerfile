FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN cd test
RUN go test
RUN cd ..
ENTRYPOINT go run server.go