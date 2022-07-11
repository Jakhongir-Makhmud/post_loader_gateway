FROM golang:1.18rc1-alpine3.15
RUN mkdir api-gateway
COPY . /api-gateway/
WORKDIR /api-gateway
RUN go build -l main cmd/main.go
CMD ./main