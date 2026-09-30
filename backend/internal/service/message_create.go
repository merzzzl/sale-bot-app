package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *Service) MessageCreate(ctx context.Context, params model.CreateMessageParams) error {
	log.Debug().Str("layer", "service").Str("func", "MessageCreate").Int64("bot_id", params.BotID).Int64("chat_id", params.ChatID).Msg("call service method")

	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		message := &model.Message{
			ChatID:     params.ChatID,
			SenderName: params.SenderName,
			SenderType: params.SenderType,
			Type:       params.Type,
			Text:       params.Text,
			ToolCalls:  params.ToolCalls,
			ToolCallID: params.ToolCallID,
		}

		if err := tx.Create(message).Error; err != nil {
			return fmt.Errorf("failed to create message: %w", err)
		}

		statistic := &model.Statistic{
			BotID: params.BotID,
			Type:  model.StatisticTypeMessages,
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
