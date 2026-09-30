package webhook

import (
	"context"
	"errors"
	"fmt"
	"unicode"

	"github.com/go-telegram/bot"
	"github.com/ltbots/backend/internal/service"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
)

var (
	ErrBotIsNotActive      = errors.New("bot is not active")
	ErrInsufficientBalance = errors.New("insufficient balance")
)

type Webhook struct {
	service      *service.Service
	bots         map[int64]func() error
	openai       *openai.Client
	cahtModel    string
	router       *Router
	messagePrice int64
	botToken     string
	secretToken  string
}

func NewWebhook(url string, service *service.Service, openai *openai.Client, chatModel string, messagePrice int64, botToken string) *Webhook {
	secretToken := make([]rune, 0, len(botToken))
	for _, r := range botToken {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			secretToken = append(secretToken, r)
		}
	}

	return &Webhook{
		service:      service,
		bots:         make(map[int64]func() error),
		openai:       openai,
		cahtModel:    chatModel,
		router:       NewRouter(url, string(secretToken)),
		messagePrice: messagePrice,
		botToken:     botToken,
		secretToken:  string(secretToken),
	}
}

func (w *Webhook) Router() *Router {
	return w.router
}

func (w *Webhook) SyncBots(ctx context.Context) error {
	if err := w.notifyBot(ctx); err != nil {
		return fmt.Errorf("failed to start notify bot: %w", err)
	}

	for ctx.Err() == nil {
		select {
		case <-ctx.Done():
			return nil
		case <-w.service.BotEvents():
			log.Info().Msg("active bot list updates")

			if errs := w.syncBots(ctx); errs != nil {
				log.Error().Errs("error", errs).Msg("sync bots error")
			}
		}
	}

	return nil
}

func (w *Webhook) notifyBot(ctx context.Context) error {
	opts := []bot.Option{
		bot.WithDefaultHandler(w.notifyBotHandler),
	}

	tgBot, err := bot.New(w.botToken, opts...)
	if err != nil {
		return fmt.Errorf("failed to create bot: %w", err)
	}

	url := w.router.add(tgBot)

	if _, err := tgBot.SetWebhook(ctx, &bot.SetWebhookParams{
		URL:         url,
		SecretToken: w.secretToken,
	}); err != nil {
		return fmt.Errorf("failed to set webhook: %w", err)
	}

	go tgBot.StartWebhook(ctx)

	go func() {
		<-ctx.Done()
		tgBot.DeleteWebhook(ctx, &bot.DeleteWebhookParams{})
	}()

	return nil
}

func (w *Webhook) syncBots(ctx context.Context) []error {
	forAdd := make(map[int64]struct{})
	forRemove := make(map[int64]struct{})

	for botID := range w.bots {
		forRemove[botID] = struct{}{}
	}

	activeBots, err := w.service.BotActiveList(ctx)
	if err != nil {
		return []error{fmt.Errorf("failed to get active bot list: %w", err)}
	}

	for _, bot := range activeBots {
		if _, ok := forAdd[bot.TelegramID]; !ok {
			forAdd[bot.TelegramID] = struct{}{}
			delete(forRemove, bot.TelegramID)
		}
	}

	errorList := make([]error, 0)

	for botID := range forRemove {
		if fn, ok := w.bots[botID]; ok {
			if err := fn(); err != nil {
				errorList = append(errorList, fmt.Errorf("stop bot %d: %w", botID, err))
			}

			log.Info().Int64("bot_id", botID).Msg("bot stopped")
		}
	}

	for botID := range forAdd {
		if err := w.startBot(ctx, botID); err != nil {
			errorList = append(errorList, fmt.Errorf("start bot %d: %w", botID, err))

			if err := w.service.BotDeactivate(ctx, botID); err != nil {
				errorList = append(errorList, fmt.Errorf("deactivate bot %d: %w", botID, err))
			}
		}

		log.Info().Int64("bot_id", botID).Msg("bot started")
	}

	if len(errorList) > 0 {
		return errorList
	}

	return nil
}

func (w *Webhook) startBot(ctx context.Context, botID int64) error {
	log.Debug().Int64("bot_id", botID).Msg("starting bot")

	b, err := w.service.BotGet(ctx, botID)
	if err != nil {
		return fmt.Errorf("failed to get bot: %w", err)
	}

	amount, err := w.service.TransactionAmount(ctx, b.UserID)
	if err != nil {
		return fmt.Errorf("failed to get user amount: %w", err)
	}

	if amount < w.messagePrice {
		return ErrInsufficientBalance
	}

	if !b.Active {
		return ErrBotIsNotActive
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(w.handler),
	}

	tgBot, err := bot.New(b.Token, opts...)
	if err != nil {
		return fmt.Errorf("failed to create bot: %w", err)
	}

	url := w.router.add(tgBot)

	if _, err := tgBot.SetWebhook(ctx, &bot.SetWebhookParams{
		URL:                url,
		DropPendingUpdates: true,
	}); err != nil {
		return fmt.Errorf("failed to set webhook: %w", err)
	}

	go tgBot.StartWebhook(ctx)

	w.bots[botID] = func() error {
		log.Debug().Int64("bot_id", botID).Msg("stopping bot")

		if _, err := tgBot.DeleteWebhook(ctx, &bot.DeleteWebhookParams{}); err != nil {
			return fmt.Errorf("failed to delete webhook: %w", err)
		}

		if _, err := tgBot.Close(ctx); err != nil {
			return fmt.Errorf("failed to close bot: %w", err)
		}

		w.router.remove(tgBot)

		return nil
	}

	return nil
}
