package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) BotActivate(ctx context.Context, botID int64) error {
	log.Debug().Str("layer", "service").Str("func", "BotActivate").Int64("bot_id", botID).Msg("call service method")

	if err := s.db.WithContext(ctx).Model(&model.Bot{}).Where("id = ?", botID).Update("active", true).Error; err != nil {
		return fmt.Errorf("failed to stop bot: %w", err)
	}

	s.botEvents <- struct{}{}

	return nil
}
