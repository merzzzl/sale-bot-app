# Backend

Backend service with OpenAI and Telegram integration for **Light Telegram Bots**

## Tech stack

- Go (service runtime)
- Telegram Bot API client (`github.com/go-telegram/bot`)
- OpenAI client (`github.com/sashabaranov/go-openai`)
- REST/Swagger from protobuf (`github.com/merzzzl/proto-rest-api`)
- Database via GORM (PostgreSQL or SQLite)

## Environment

- `APP_HOSTNAME` (default: `localhost`)
- `APP_OPENAI_API_KEY` (required)
- `APP_OPENAI_MODEL` (default: `gpt-5-mini`)
- `APP_MAIN_BOT_TOKEN` (required)
- `APP_DB_DRIVER` (default: `sqlite`, supported: `sqlite`, `postgres`)
- `APP_DB_URL` (default: `sqlite.db`)
- `APP_MESSAGE_PRICE` (default: `50`)

## Related repositories

- Contracts / API clients: https://github.com/ltbots/protocols
- Example docker-compose for the whole project: https://github.com/ltbots/compose