FROM golang:1.19-alpine as builder

WORKDIR /consumer-microservice/

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o microservice /consumer-microservice/main.go

FROM alpine:latest

RUN apk add --no-cache libc6-compat


WORKDIR /consumer-microservice

COPY --from=builder /consumer-microservice/ /consumer-microservice/

EXPOSE 9090

CMD ./microservice
