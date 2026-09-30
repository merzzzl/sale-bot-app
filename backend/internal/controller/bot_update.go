package controller

import (
	"context"
	"net/http"

	"github.com/ltbots/backend/internal/model"
	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
)

func (c *ControllerService) BotUpdate(ctx context.Context, req *api.BotUpdateRequest) (*api.Bot, error) {
	if err := c.access(ctx, accessCheck{botID: req.GetBotId()}); err != nil {
		log.Error().Err(err).Msg("failed to check bot access")

		return nil, runtime.Error(http.StatusForbidden, "failed to check bot access")
	}

	if err := c.service.BotUpdate(ctx, req.GetBotId(), model.UpdateBotParams{
		Prompt:   req.GetPrompt(),
		PresetID: req.GetPresetId(),
	}); err != nil {
		log.Error().Err(err).Msg("failed to update bot")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to update bot")
	}

	bot, err := c.service.BotGet(ctx, req.GetBotId())
	if err != nil {
		log.Error().Err(err).Msg("failed to get bot after update")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to get bot after update")
	}

	return &api.Bot{
		BotId:    bot.TelegramID,
		Prompt:   bot.Prompt,
		Active:   bot.Active,
		Token:    bot.Token,
		PresetId: bot.PresetID,
	}, nil
}
