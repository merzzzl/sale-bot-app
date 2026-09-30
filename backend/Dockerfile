# Build with the monorepo root as context: docker build -f backend/Dockerfile .
FROM golang:1.26-alpine AS builder
WORKDIR /app/backend
RUN apk add --no-cache gcc musl-dev
COPY protocols/ /app/protocols/
COPY backend/ ./
ENV CGO_ENABLED=1
RUN go build -trimpath -o /out/sale-bot-app ./cmd/app
FROM alpine:3.23
RUN apk add --no-cache ca-certificates libgcc
WORKDIR /app
COPY --from=builder /out/sale-bot-app /app/sale-bot-app
EXPOSE 8080
CMD ["/app/sale-bot-app"]
