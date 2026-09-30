package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) BotAccess(ctx context.Context, userID, botID int64) error {
	log.Debug().Str("layer", "service").Str("func", "BotAccess").Int64("user_id", userID).Int64("bot_id", botID).Msg("call service method")

	if err := s.db.WithContext(ctx).Where("user_id = ? AND id = ?", userID, botID).First(&model.Bot{}).Error; err != nil {
		return fmt.Errorf("failed to get bot: %w", err)
	}

	return nil
}
