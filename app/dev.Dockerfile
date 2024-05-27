# Choose whatever you want, version >= 1.16
FROM golang:1.22-alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app


# COPY go.mod go.sum ./
# RUN go mod download
ENTRYPOINT ["air"]

# CMD ["air", "-c", ".air.toml"]