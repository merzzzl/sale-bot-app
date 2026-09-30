package payment

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/ltbots/backend/internal/i18n"
	"github.com/rs/zerolog/log"
)

func (a *Agent) SuccessfulPayment(ctx context.Context) (bool, error) {
	if a.update.Message == nil || a.update.Message.SuccessfulPayment == nil {
		return false, nil
	}

	b, err := a.service.BotGet(ctx, a.bot.ID())
	if err != nil {
		return false, fmt.Errorf("failed to get bot: %w", err)
	}

	user, err := a.service.UserGet(ctx, b.UserID)
	if err != nil {
		return false, fmt.Errorf("failed to get user: %w", err)
	}

	payload := struct {
		ProductID int64  `json:"product_id"`
		Timestamp string `json:"timestamp"`
	}{}

	if err := json.Unmarshal([]byte(a.update.Message.SuccessfulPayment.InvoicePayload), &payload); err != nil {
		return false, fmt.Errorf("failed to unmarshal invoice payload: %w", err)
	}

	product, err := a.service.ProductGet(ctx, payload.ProductID)
	if err != nil {
		return false, fmt.Errorf("failed to get product: %w", err)
	}

	if err := a.service.StatisticsPayment(ctx, a.bot.ID()); err != nil {
		return false, fmt.Errorf("failed to create payment statistic: %w", err)
	}

	if _, err := a.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: a.update.Message.Chat.ID,
		Text:   i18n.Localize(user.LanguageCode, "notify_success_message"),
	}); err != nil {
		log.Error().Err(err).Msg("failed to send message")
	}

	if user.NotificationChatID == 0 {
		return true, nil
	}

	if _, err := a.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: user.NotificationChatID,
		Text: fmt.Sprintf(i18n.Localize(user.LanguageCode, "payment_notify_message"),
			product.Name,
			product.Price,
			product.Currency,
			a.update.Message.From.Username,
		),
	}); err != nil {
		log.Error().Err(err).Msg("failed to send message")
	}

	return true, nil
}
