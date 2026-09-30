# sale-bot-app

**sale-bot-app** is a Telegram Mini App platform for running AI-assisted sales
bots: connect a bot, configure its agent, manage products/services, and accept
payments inside Telegram.

## What’s inside

- **[Frontend](frontend)** — Telegram Mini App UI (React, Chakra UI, Vite).
- **[Backend](backend)** — AI and Telegram integrations (Go, GORM, SQLite/Postgres).
- **[Contracts / API clients](protocols)** — protobuf, Swagger and Go/TS clients.
- **[Deployment examples](compose)** — Docker Compose for the whole stack.

## Architecture

- The **Mini App UI** talks to the **backend API** on the same origin.
- The **backend** integrates with Telegram (bot/webhooks), uses an LLM provider
  for agent responses, and stores data in a database.
- The API surface is defined in **protobuf contracts** and shared across clients.
- One Docker image, `ghcr.io/merzzzl/sale-bot-app`, contains the Go backend and
  compiled Mini App. Caddy provides HTTPS in the Compose deployment.

## Repository map

This monorepo keeps focused components in separate directories:

```text
backend/     Go application and integrations
frontend/    Telegram Mini App UI
protocols/   Contracts and generated Go/TypeScript API clients
compose/     Deployment configuration
screenshot/ Mini App screenshots
```

The backend uses the local `protocols` Go module through a `replace` directive.
The frontend uses `@sale-bot-app/api` through `file:../protocols`.
No separate API-client release is needed when contracts change.

## Quick start

```sh
git clone https://github.com/merzzzl/sale-bot-app.git
cd sale-bot-app/compose
cp .env.example .env
# Fill in your domain, Telegram bot token and OpenAI API key.
docker compose up -d --build
```

See [deployment instructions](compose/README.md) for DNS, HTTPS and persistence.
Configure the main bot's Mini App URL in BotFather to point to your HTTPS domain.

## Development

Requirements: Go 1.26+, Node.js 24+, npm, and a C compiler for SQLite.

```sh
make install
make build
make test
make lint
```

Run the backend from `backend/` with `go run ./cmd/app` after setting the
[environment variables](backend/README.md). It serves the built frontend from
`../frontend/dist` (override with `APP_STATIC_DIR`). For frontend hot reload,
run `npm --prefix frontend run dev`; Vite proxies API/webhook/Telegram requests
to `localhost:8080`.

Regenerate contracts with `make -C protocols generate` (requires protoc,
protoc-gen-go, protoc-gen-go-rest, Docker, and npm).
For rootless Docker, pass `DOCKER_USER=0:0` when regenerating the TS client.

## Container publishing

The Docker workflow publishes `ghcr.io/merzzzl/sale-bot-app` when a GitHub release
is published, or on manual dispatch. Each build is also tagged `latest`.
CI checks Go tests, lint, frontend/client builds and the combined Docker image.

TypeScript is kept on 6.0.3, the latest version supported by typescript-eslint.
Existing React Compiler form-state/ref diagnostics are reported as warnings;
frontend lint has no errors. Vite also warns about the existing large UI bundle.

Dependency checks: both npm audits report no known vulnerabilities. Go's
govulncheck reports no reachable vulnerable symbols. It flags GO-2026-6443 in
the indirect gRPC module, but this app does not import the affected server
package; the available fix is currently a development version rather than
a stable release.

## Screenshots

<p>
  <img src="screenshot/img1.png" width="25%" alt="Screenshot 1">
  <img src="screenshot/img2.png" width="25%" alt="Screenshot 2">
  <img src="screenshot/img3.png" width="25%" alt="Screenshot 3">
</p>
