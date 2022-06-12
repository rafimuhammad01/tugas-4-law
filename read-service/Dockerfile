FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download

COPY . /app

RUN go build -o /read-service

CMD [ "/read-service" ]

FROM alpine:latest

COPY --from=builder /read-service .

CMD [ "./read-service" ]