package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *Service) ChatCreate(ctx context.Context, telegramID int64, botID int64) error {
	log.Debug().Str("layer", "service").Str("func", "ChatCreate").Int64("telegram_id", telegramID).Int64("bot_id", botID).Msg("call service method")

	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var chat model.Chat

		err := tx.Find(&chat, "id = ? AND bot_id = ?", telegramID, botID).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed to find chat: %w", err)
		}

		if chat.TelegramID == telegramID {
			return nil
		}

		chat = model.Chat{
			TelegramModel: model.TelegramModel{
				TelegramID: telegramID,
			},
			BotID: botID,
		}

		if err := tx.Create(&chat).Error; err != nil {
			return fmt.Errorf("failed to create chat: %w", err)
		}

		statistic := &model.Statistic{
			BotID: botID,
			Type:  model.StatisticTypeChats,
			Value: 1,
		}

		if err := tx.Create(statistic).Error; err != nil {
			return fmt.Errorf("failed to create statistic: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	return nil
}
