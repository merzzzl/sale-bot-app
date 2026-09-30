module github.com/merzzzl/sale-bot-app/backend

go 1.26.0

replace github.com/merzzzl/sale-bot-app/protocols => ../protocols

require (
	github.com/caarlos0/env/v11 v11.4.1
	github.com/go-telegram/bot v1.27.0
	github.com/merzzzl/proto-rest-api v0.0.1-alpha.29
	github.com/merzzzl/sale-bot-app/protocols v0.0.0
	github.com/rs/zerolog v1.35.1
	github.com/sashabaranov/go-openai v1.43.0
	github.com/telegram-mini-apps/init-data-golang v1.5.0
	google.golang.org/protobuf v1.36.12
	gorm.io/driver/postgres v1.6.3
	gorm.io/driver/sqlite v1.6.0
	gorm.io/gorm v1.31.2
)

require (
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.11.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/mattn/go-colorable v0.1.15 // indirect
	github.com/mattn/go-isatty v0.0.24 // indirect
	github.com/mattn/go-sqlite3 v1.14.52 // indirect
	github.com/swaggo/files/v2 v2.0.2 // indirect
	golang.org/x/sync v0.23.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
	golang.org/x/text v0.42.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260928230214-8a89bd6388cc // indirect
	google.golang.org/grpc v1.84.0 // indirect
)
