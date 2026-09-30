package controller

import (
	"context"
	"net/http"

	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *ControllerService) BotList(ctx context.Context, req *emptypb.Empty) (*api.BotListResponse, error) {
	initData, err := GetInitDataFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get init data")

		return nil, runtime.Error(http.StatusUnauthorized, "failed to get init data")
	}

	bots, err := c.service.BotList(ctx, initData.User.ID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get bot list")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to get bot list")
	}

	var response api.BotListResponse

	for _, bot := range bots {
		response.Bots = append(response.Bots, &api.Bot{
			BotId:    bot.TelegramID,
			Prompt:   bot.Prompt,
			Active:   bot.Active,
			Token:    bot.Token,
			PresetId: bot.PresetID,
		})
	}

	return &response, nil
}
