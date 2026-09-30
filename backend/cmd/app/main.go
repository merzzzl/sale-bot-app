package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-telegram/bot"
	"github.com/ltbots/backend/internal/controller"
	"github.com/ltbots/backend/internal/model"
	"github.com/ltbots/backend/internal/service"
	"github.com/ltbots/backend/internal/webhook"
	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/merzzzl/proto-rest-api/swagger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("reading config...")

	config, err := readConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read config")
	}

	var dialector gorm.Dialector

	switch config.DBDriver {
	case "postgres":
		dialector = postgres.Open(config.DBURL)
	case "sqlite":
		dialector = sqlite.Open(config.DBURL)
	}

	log.Info().Msg("connect to database...")

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open database")
	}

	log.Info().Msg("migrating database...")

	if err := db.AutoMigrate(
		&model.Bot{},
		&model.Chat{},
		&model.Message{},
		&model.Product{},
		&model.Statistic{},
		&model.Transaction{},
		&model.User{},
	); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate database")
	}

	log.Info().Msg("connect to telegram...")

	mainBot, err := bot.New(config.MainBotToken)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to telegram")
	}

	log.Info().Msg("connect to openai...")

	openaiClient := openai.NewClient(config.OpenAIAPIKey)

	if _, err := openaiClient.ListModels(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to connect to openai")
	}

	webhookURL := fmt.Sprintf("https://%s/", config.Hostname)

	ser := service.NewService(db, mainBot)
	ctrl := controller.NewController(ser, config.MainBotToken)
	wh := webhook.NewWebhook(webhookURL, ser, openaiClient, config.OpenAIModel, config.MessagePrice, config.MainBotToken)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := ser.InitDB(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to init db presets")
	}

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	go func() {
		if err := wh.SyncBots(ctx); err != nil {
			log.Fatal().Err(err).Msg("failed to sync bots")
		}
	}()

	router := runtime.NewRouter()

	api.RegisterControllerHandler(router, ctrl, ctrl.InitDataMiddleware(), ctrl.LogMiddleware())

	mux := router.Mux()

	mux.Handle("/api/v1/swagger-ui/", swagger.Handler(api.GetV1Swagger()))
	mux.Handle("/webhook/", wh.Router())

	router.Router.HandleMethodNotAllowed = true
	router.Router.HandleOPTIONS = true
	router.Router.RedirectFixedPath = true
	router.Router.RedirectTrailingSlash = true

	httpServer := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", config.Hostname)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			mux.ServeHTTP(w, r)
		}),
	}

	log.Info().Msg("serving http...")

	go func() {
		if err := httpServer.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("failed to serve")
		}
	}()

	log.Info().Msg("server is running on :8080")

	signals := make(chan os.Signal, 1)

	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals

	httpServer.Shutdown(ctx)
	lis.Close()
	cancel()

	log.Info().Msg("server is stopped")
}
