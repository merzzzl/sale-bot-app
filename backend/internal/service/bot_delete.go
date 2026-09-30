package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) BotDelete(ctx context.Context, botID int64) error {
	log.Debug().Str("layer", "service").Str("func", "BotDelete").Int64("bot_id", botID).Msg("call service method")

	if err := s.db.WithContext(ctx).Delete(&model.Bot{}, botID).Error; err != nil {
		return fmt.Errorf("failed to delete bot: %w", err)
	}

	s.botEvents <- struct{}{}

	return nil
}
