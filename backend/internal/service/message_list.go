package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) MessageList(ctx context.Context, chatID int64) ([]model.Message, error) {
	log.Debug().Str("layer", "service").Str("func", "MessageList").Int64("chat_id", chatID).Msg("call service method")

	var messages []model.Message

	if err := s.db.WithContext(ctx).Where("chat_id = ?", chatID).Order("id ASC").Find(&messages).Error; err != nil {
		return nil, fmt.Errorf("не удалось получить список сообщений: %w", err)
	}

	return messages, nil
}
