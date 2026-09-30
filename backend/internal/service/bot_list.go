package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) BotList(ctx context.Context, UserID int64) ([]model.Bot, error) {
	log.Debug().Str("layer", "service").Str("func", "BotList").Int64("user_id", UserID).Msg("call service method")

	var bots []model.Bot

	if err := s.db.WithContext(ctx).Where("user_id = ?", UserID).Find(&bots).Error; err != nil {
		return nil, fmt.Errorf("failed to get bot list: %w", err)
	}

	return bots, nil
}
