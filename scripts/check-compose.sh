#!/bin/sh
set -eu

# Only fixture values: never load the developer's .env or contact services.
export APP_HOSTNAME=sale-bot.example.com APP_OPENAI_API_KEY=ci-placeholder
export APP_MAIN_BOT_TOKEN=ci-placeholder ACME_EMAIL=admin@example.com
export MINI_APP_DOMAIN=app.example.com WG_EASY_DOMAIN=wg.example.com
export WG_EASY_USERNAME=admin WG_EASY_PASSWORD=ci-placeholder
export WG_PUBLIC_HOST=wg.example.com

find compose example -type f \( -name compose.yaml -o -name docker-compose.yaml \) -print | while IFS= read -r file; do
  printf 'Checking %s\n' "$file"
  docker compose --env-file "$(dirname "$file")/.env.example" -f "$file" config --quiet
done
