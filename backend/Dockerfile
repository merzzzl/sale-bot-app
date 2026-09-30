FROM golang:1.24.6-alpine AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=1
RUN apk add --no-cache gcc-go musl-dev
RUN go build -o app cmd/app/*

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app /app/app
EXPOSE 8080
CMD ["./app"]
