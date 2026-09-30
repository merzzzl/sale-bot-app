#!/bin/sh
set -eu
umask 077

REPOSITORY="${REPOSITORY:-merzzzl/sale-bot-app}"
DEPLOY_REF="${DEPLOY_REF:-main}"
RAW_BASE_URL="${RAW_BASE_URL:-https://raw.githubusercontent.com/${REPOSITORY}/${DEPLOY_REF}}"
FORCE=0

case "${1:-}" in
  '') ;;
  --force) FORCE=1 ;;
  *) printf 'Usage: sh bootstrap-deploy.sh [--force]\n' >&2; exit 2 ;;
esac
[ "$#" -le 1 ] || { printf 'Too many arguments\n' >&2; exit 2; }

for file in compose.yaml Caddyfile .env.example README.md; do
  if [ -e "$file" ] && [ "$FORCE" -eq 0 ]; then
    printf '%s exists; use --force to update deployment files. .env is always preserved.\n' "$file" >&2
    exit 1
  fi
done

stage=$(mktemp -d)
trap 'rm -rf "$stage"' EXIT HUP INT TERM

download_file() {
  source_url="${RAW_BASE_URL}/example/$1"
  if command -v curl >/dev/null 2>&1; then
    curl -fsSL "$source_url" -o "$stage/$1"
  elif command -v wget >/dev/null 2>&1; then
    wget -qO "$stage/$1" "$source_url"
  else
    printf 'curl or wget is required\n' >&2
    exit 1
  fi
}

# Download everything before modifying any existing deployment files.
for file in compose.yaml Caddyfile .env.example README.md; do
  download_file "$file"
done
for file in compose.yaml Caddyfile .env.example README.md; do
  cp "$stage/$file" "$file"
done

if [ ! -e .env ]; then
  cp .env.example .env
  chmod 600 .env
  printf 'Created .env. Fill in the required values before starting.\n'
else
  printf 'Preserved existing .env. Review .env.example for new variables.\n'
fi
printf 'Next: edit .env, then run docker compose pull && docker compose up -d\n'
