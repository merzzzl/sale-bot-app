package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ltbots/backend/internal/i18n"
	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) TransactionBill(ctx context.Context, userID int64, amount int64) error {
	log.Debug().Str("layer", "service").Str("func", "TransactionBill").Int64("user_id", userID).Int64("amount", amount).Msg("call service method")

	user := &model.User{}

	if err := s.db.WithContext(ctx).First(user, userID).Error; err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}

	if user.NotificationChatID == 0 {
		return fmt.Errorf("user has no notification chat")
	}

	if _, err := s.mainBot.SendInvoice(ctx, &bot.SendInvoiceParams{
		ChatID:      user.NotificationChatID,
		Title:       i18n.Localize(user.LanguageCode, "send_invoice_title"),
		Description: fmt.Sprintf(i18n.Localize(user.LanguageCode, "send_invoice_description"), amount/100),
		Payload:     time.Now().Format(time.RFC3339),
		Currency:    "XTR",
		Prices: []models.LabeledPrice{
			{
				Label:  i18n.Localize(user.LanguageCode, "send_invoice_price_label"),
				Amount: int(amount / 100),
			},
		},
	}); err != nil {
		return fmt.Errorf("failed to send invoice: %w", err)
	}

	return nil
}
