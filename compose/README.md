# sale-bot-app deployment

The backend, Mini App UI and API clients are built from this monorepo into
`ghcr.io/merzzzl/sale-bot-app:latest`. Caddy provides HTTPS for Telegram webhooks.

```sh
git clone https://github.com/merzzzl/sale-bot-app.git
cd sale-bot-app/compose
cp .env.example .env
# Set your public hostname, ACME email, Telegram token and OpenAI key in .env.
docker compose up -d --build
```

Point the hostname's DNS record to your server and open ports 80 and 443.
Configure the main bot's Mini App URL as `https://<APP_HOSTNAME>` in BotFather.
SQLite and Caddy certificates are stored in persistent Docker volumes.

To use a published image instead of building locally:

```sh
docker compose pull
docker compose up -d --no-build
```

Do not commit `.env` or database files. Back up the `sale-bot-app_data` volume.

## Configuration reference

<!-- BEGIN:compose.yaml -->
```yaml
services:
  sale-bot-app:
    image: ghcr.io/merzzzl/sale-bot-app:latest
    build:
      context: ..
    restart: unless-stopped
    environment:
      APP_HOSTNAME: ${APP_HOSTNAME:?Set APP_HOSTNAME in .env}
      APP_OPENAI_API_KEY: ${APP_OPENAI_API_KEY:?Set APP_OPENAI_API_KEY in .env}
      APP_OPENAI_MODEL: ${APP_OPENAI_MODEL:-gpt-5-mini}
      APP_MAIN_BOT_TOKEN: ${APP_MAIN_BOT_TOKEN:?Set APP_MAIN_BOT_TOKEN in .env}
      APP_DB_DRIVER: sqlite
      APP_DB_URL: /data/sqlite.db
      APP_MESSAGE_PRICE: ${APP_MESSAGE_PRICE:-50}
    volumes:
      - sale-bot-app_data:/data

  caddy:
    image: caddy:2.11-alpine
    restart: unless-stopped
    depends_on:
      - sale-bot-app
    environment:
      APP_HOSTNAME: ${APP_HOSTNAME}
      ACME_EMAIL: ${ACME_EMAIL}
    ports:
      - "80:80"
      - "443:443"
      - "443:443/udp"
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile:ro
      - caddy_data:/data
      - caddy_config:/config

volumes:
  sale-bot-app_data:
  caddy_data:
  caddy_config:
```
<!-- END:compose.yaml -->

<!-- BEGIN:Caddyfile -->
```caddyfile
{
  email {$ACME_EMAIL}
}

{$APP_HOSTNAME} {
  reverse_proxy sale-bot-app:8080
}
```
<!-- END:Caddyfile -->
