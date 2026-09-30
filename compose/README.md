# Compose

Docker Compose stack for running the services locally using prebuilt container images.

## What’s inside

This stack runs three services:

- **backend**: `ghcr.io/ltbots/backend:latest`
- **frontend**: `ghcr.io/ltbots/frontend:latest`
- **caddy**: `caddy:2.11-alpine`

Caddy routing:

- `/api/*` → `backend:8080`
- `/webhook/*` → `backend:8080`
- `/tg/*` → Telegram API (`https://api.telegram.org`)
- everything else → `frontend:8080`

## Example configs
compose.yaml:
<!-- BEGIN:compose.yaml -->
```yaml
services:
  backend:
    image: ghcr.io/ltbots/backend:latest
    restart: unless-stopped
    environment:
      APP_HOSTNAME: localhost
      APP_OPENAI_API_KEY: <past you open-ai key>
      APP_OPENAI_MODEL: gpt-5-mini
      APP_MAIN_BOT_TOKEN: <past you bot token>
      APP_DB_DRIVER: sqlite
      APP_DB_URL: /data/sqlite.db
      APP_MESSAGE_PRICE: 50

  frontend:
    image: ghcr.io/ltbots/frontend:latest
    restart: unless-stopped

  caddy:
    image: caddy:2.11-alpine
    restart: unless-stopped
    depends_on:
      - backend
      - frontend
    ports:
      - "8080:80"
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile:ro
```
<!-- END:compose.yaml -->

Caddyfile:
<!-- BEGIN:Caddyfile -->
```caddyfile
:80 {
  handle_path /tg/* {
    reverse_proxy https://api.telegram.org {
      header_up Host api.telegram.org
    }
  }

  handle /api/* {
    reverse_proxy backend:8080
  }

  handle /webhook/* {
    reverse_proxy backend:8080
  }

  handle {
    reverse_proxy frontend:8080
  }
}
```
<!-- END:Caddyfile -->

## Quick start

```sh
git clone https://github.com/ltbots/compose
cd compose
docker compose up -d
```
