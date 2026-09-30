package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) BotActiveList(ctx context.Context) ([]model.Bot, error) {
	log.Debug().Str("layer", "service").Str("func", "BotActiveList").Msg("call service method")

	var bots []model.Bot

	if err := s.db.WithContext(ctx).Where("active = true AND prompt != ''").Find(&bots).Error; err != nil {
		return nil, fmt.Errorf("failed to get bot list: %w", err)
	}

	return bots, nil
}
