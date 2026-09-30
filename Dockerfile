FROM golang:1.26-alpine AS backend
WORKDIR /app/backend
RUN apk add --no-cache gcc musl-dev
COPY protocols/go.mod protocols/go.sum /app/protocols/
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY protocols/ /app/protocols/
COPY backend/ ./
ENV CGO_ENABLED=1
RUN go build -trimpath -o /out/sale-bot-app ./cmd/app

FROM node:24-alpine AS frontend
WORKDIR /app
COPY protocols/ ./protocols/
COPY frontend/package*.json ./frontend/
RUN cd protocols && npm ci
RUN cd protocols && npm run build
RUN cd frontend && npm ci
COPY frontend/ ./frontend/
RUN cd frontend && npm run build

FROM alpine:3.23
RUN apk add --no-cache ca-certificates libgcc
WORKDIR /app
COPY --from=backend /out/sale-bot-app /app/sale-bot-app
COPY --from=frontend /app/frontend/dist /app/static
ENV APP_STATIC_DIR=/app/static
EXPOSE 8080
CMD ["/app/sale-bot-app"]
