FROM golang:1.19-alpine

ARG SERVER_PORT

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/payment-manager main.go

EXPOSE ${SERVER_PORT}

ENTRYPOINT ["/app/payment-manager"]
