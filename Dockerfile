FROM golang:1.16-alpine AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o vault-dotenv main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/vault-dotenv /usr/local/bin/.
CMD ["vault-dotenv"]
