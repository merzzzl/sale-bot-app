package service

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) BotCreate(ctx context.Context, params model.CreateBotParams) (*model.Bot, error) {
	log.Debug().Str("layer", "service").Str("func", "BotCreate").Int64("user_id", params.UserID).Msg("call service method")

	tg, err := bot.New(params.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	bot := &model.Bot{
		TelegramModel: model.TelegramModel{
			TelegramID: tg.ID(),
		},
		Token:    params.Token,
		UserID:   params.UserID,
		PresetID: params.PresetID,
	}

	if err := s.db.WithContext(ctx).Create(bot).Error; err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	return bot, nil
}
