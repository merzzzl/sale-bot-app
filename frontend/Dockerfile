# Build with the monorepo root as context: docker build -f frontend/Dockerfile .
FROM node:24-alpine AS builder
WORKDIR /app
COPY protocols/ ./protocols/
COPY frontend/ ./frontend/
RUN cd protocols && npm ci && npm run build
RUN cd frontend && npm ci && npm run build
FROM nginx:1.29-alpine
COPY --from=builder /app/frontend/dist /usr/share/nginx/html
COPY frontend/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080
