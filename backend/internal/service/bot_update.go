package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *Service) BotUpdate(ctx context.Context, botID int64, params model.UpdateBotParams) error {
	log.Debug().Str("layer", "service").Str("func", "BotUpdate").Int64("bot_id", botID).Msg("call service method")

	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var bot model.Bot

		if err := tx.First(&bot, botID).Error; err != nil {
			return fmt.Errorf("failed to get bot: %w", err)
		}

		updateData := make(map[string]interface{})

		if params.Prompt != "" && params.Prompt != bot.Prompt {
			updateData["prompt"] = params.Prompt
		}

		if params.PresetID != 0 && params.PresetID != bot.PresetID {
			updateData["preset_id"] = params.PresetID
		}

		if err := tx.Model(&bot).Where("id = ?", botID).Updates(updateData).Error; err != nil {
			return fmt.Errorf("failed to update bot: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	return nil
}
