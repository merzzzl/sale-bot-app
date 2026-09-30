package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) BotGet(ctx context.Context, botID int64) (*model.Bot, error) {
	log.Debug().Str("layer", "service").Str("func", "BotGet").Int64("bot_id", botID).Msg("call service method")

	var bot *model.Bot

	if err := s.db.WithContext(ctx).Find(&bot, botID).Error; err != nil {
		return nil, fmt.Errorf("failed to get bot: %w", err)
	}

	return bot, nil
}
