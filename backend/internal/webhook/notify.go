package webhook

import (
	"context"
	"errors"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ltbots/backend/internal/i18n"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (w *Webhook) notifyBotHandler(ctx context.Context, tgBot *bot.Bot, update *models.Update) {
	log.Debug().Str("layer", "webhook").Str("worker", "notify").Msg("call webhook method")

	if update.PreCheckoutQuery != nil {
		if _, err := tgBot.AnswerPreCheckoutQuery(ctx, &bot.AnswerPreCheckoutQueryParams{
			PreCheckoutQueryID: update.PreCheckoutQuery.ID,
			OK:                 true,
		}); err != nil {
			log.Error().Err(err).Msg("failed to answer pre checkout query")

			return
		}

		return
	}

	if update.Message == nil {
		return
	}

	user, err := w.service.UserGet(ctx, update.Message.From.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).Msg("failed to get user")

		return
	}

	if user == nil || user.NotificationChatID != update.Message.Chat.ID || update.Message.Text == "/start" {
		if err := w.service.UserCreate(ctx, update.Message.From.ID, update.Message.Chat.ID, user.LanguageCode); err != nil {
			log.Error().Err(err).Msg("failed to create user")

			return
		}

		if _, err := tgBot.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   i18n.Localize(user.LanguageCode, "notify_message"),
		}); err != nil {
			log.Error().Err(err).Msg("failed to send message")

			return
		}

		return
	}

	if update.Message != nil && update.Message.SuccessfulPayment != nil {
		amount := int64(update.Message.SuccessfulPayment.TotalAmount * 100)

		if err := w.service.TransactionDeposit(ctx, amount, update.Message.From.ID); err != nil {
			log.Error().Err(err).Msg("failed to create transaction")

			return
		}

		if _, err := tgBot.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   i18n.Localize(user.LanguageCode, "notify_success_message"),
		}); err != nil {
			log.Error().Err(err).Msg("failed to send message")
		}

		return
	}

	if _, err := tgBot.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.Message.Chat.ID,
		MessageID: update.Message.ID,
	}); err != nil {
		log.Error().Err(err).Msg("failed to delete message")

		return
	}
}
