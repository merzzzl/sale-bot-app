package controller

import (
	"context"
	"net/http"

	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
)

func (c *ControllerService) BotGet(ctx context.Context, req *api.BotGetRequest) (*api.Bot, error) {
	if err := c.access(ctx, accessCheck{botID: req.GetBotId()}); err != nil {
		log.Error().Err(err).Msg("failed to check bot access")

		return nil, runtime.Error(http.StatusForbidden, "failed to check bot access")
	}

	bot, err := c.service.BotGet(ctx, req.GetBotId())
	if err != nil {
		log.Error().Err(err).Msg("failed to get bot")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to get bot")
	}

	return &api.Bot{
		BotId:    bot.TelegramID,
		Prompt:   bot.Prompt,
		Active:   bot.Active,
		Token:    bot.Token,
		PresetId: bot.PresetID,
	}, nil
}
