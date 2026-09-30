# Standalone sale-bot-app deployment

This example uses the published `ghcr.io/merzzzl/sale-bot-app` image.
Only Docker Engine and Docker Compose are needed; no Go, Node.js or source
checkout is required. The image must have been published by the Publish workflow.

## Installation

Copy `compose.yaml`, `Caddyfile` and `.env.example` to your deployment directory,
or use the [installer](../scripts/bootstrap-deploy.sh):

```sh
mkdir -p sale-bot-app && cd sale-bot-app
curl -fsSL https://raw.githubusercontent.com/merzzzl/sale-bot-app/main/scripts/bootstrap-deploy.sh -o bootstrap-deploy.sh
# Review the downloaded script before running it.
sh bootstrap-deploy.sh
```

If copying files manually, run `cp .env.example .env` and `chmod 600 .env`.
Edit `.env`: set the public hostname (without https://), ACME email, Telegram
bot token and OpenAI API key. Point DNS at the server and open ports 80 and 443.
Configure the main bot's Mini App URL in BotFather as `https://<APP_HOSTNAME>`.

```sh
docker compose pull
docker compose up -d
docker compose logs -f sale-bot-app
```

The installer never overwrites an existing `.env`. Updating other existing files
requires `sh bootstrap-deploy.sh --force`; review new `.env.example` variables.
Set `DEPLOY_REF=vX.Y.Z` to download example files from a specific release tag.

## Updates and rollback

Set `APP_IMAGE` in `.env` to a published version tag for reproducible deployments.
Use `latest` only if you intentionally want the most recent stable image.

```sh
docker compose pull
docker compose up -d
```

To roll back, restore the prior `APP_IMAGE` and repeat these commands.
Application database compatibility still needs to be checked before rollback.

## Persistence and backup

Named volumes store SQLite (`sale-bot-app_data`) and Caddy certificates/config.
Do not run `docker compose down -v` unless you intend to delete that data.
For a consistent database backup stop the application first:

```sh
mkdir -p backups
docker compose stop sale-bot-app
docker compose cp sale-bot-app:/data/. ./backups/
docker compose start sale-bot-app
```

Protect backups and `.env`; they can contain personal data and credentials.
Also back up the Caddy volumes using your Docker volume backup tooling.
