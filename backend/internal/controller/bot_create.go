package controller

import (
	"context"
	"net/http"

	"github.com/ltbots/backend/internal/model"
	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
)

func (c *ControllerService) BotCreate(ctx context.Context, req *api.BotCreateRequest) (*api.Bot, error) {
	initData, err := GetInitDataFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get init data")

		return nil, runtime.Error(http.StatusUnauthorized, "failed to get init data")
	}

	bot, err := c.service.BotCreate(ctx, model.CreateBotParams{
		Token:    req.GetBotToken(),
		UserID:   initData.User.ID,
		PresetID: 1,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create bot")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to create bot")
	}

	return &api.Bot{
		BotId:    bot.TelegramID,
		Prompt:   bot.Prompt,
		Active:   bot.Active,
		Token:    bot.Token,
		PresetId: bot.PresetID,
	}, nil
}
