FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download

COPY . /app

RUN go build -o /api-gateway

CMD [ "/api-gateway" ]

FROM alpine:latest

COPY --from=builder /api-gateway .

CMD [ "./api-gateway" ]

