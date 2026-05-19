FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o stress-test .

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/stress-test .
ENTRYPOINT ["./stress-test"]
