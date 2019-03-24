FROM golang:alpine

RUN apk add --no-cache git build-base

WORKDIR ./ciklum_test

COPY go.mod .
COPY go.sum .
COPY ports.json .

COPY protobuf/ ./protobuf
COPY ./ClientAPI ./ClientAPI

RUN go mod download
RUN go test ./ClientAPI/...
RUN go build -o ./ClientAPI/main ./ClientAPI/main.go

EXPOSE 5000
ENTRYPOINT ["./ClientAPI/main"]