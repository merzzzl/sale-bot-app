# Development deployment

This directory runs the application built from the monorepo root, together with
Caddy. For deployment without source code use [../example](../example).

```sh
cp .env.example .env
# Configure hostname, ACME email, Telegram token and OpenAI key.
docker compose up -d --build
```

Point the configured hostname at the server and allow TCP ports 80 and 443.
Set the main bot's Mini App URL in BotFather to `https://<APP_HOSTNAME>`.
The public HTTPS address is also used for managed bots' Telegram webhooks.

SQLite and Caddy certificates use named Docker volumes. Do not commit `.env`.
See the [standalone example](../example/README.md) for updates and backups.
