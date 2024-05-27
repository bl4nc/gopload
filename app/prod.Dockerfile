# Use a base image suitable for production
FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main .

# Use a minimal base image for running the application
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main

CMD ["/app/main"]
