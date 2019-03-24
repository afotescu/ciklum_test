FROM golang:alpine

RUN apk add --no-cache git

WORKDIR ./ciklum_test

COPY go.mod .
COPY go.sum .

COPY protobuf/ ./protobuf
COPY ./PortDomainService ./PortDomainService

RUN go mod download
RUN go test ./PortDomainService/...
RUN go build -o ./PortDomainService/main ./PortDomainService/main.go

EXPOSE 5001
ENTRYPOINT ["./PortDomainService/main"]