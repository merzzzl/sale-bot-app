package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) PromptPresetList(ctx context.Context) ([]model.PromptPreset, error) {
	log.Debug().Str("layer", "service").Str("func", "PromptPresetList").Msg("call service method")

	var presets []model.PromptPreset

	if err := s.db.WithContext(ctx).Find(&presets).Error; err != nil {
		return nil, fmt.Errorf("failed to get prompt presets: %w", err)
	}

	return presets, nil
}
