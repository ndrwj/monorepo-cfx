FROM golang:1.25.3 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
