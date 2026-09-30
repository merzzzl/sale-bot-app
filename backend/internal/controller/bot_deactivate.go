package controller

import (
	"context"
	"net/http"

	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *ControllerService) BotDeactivate(ctx context.Context, req *api.BotDeactivateRequest) (*emptypb.Empty, error) {
	if err := c.access(ctx, accessCheck{botID: req.GetBotId()}); err != nil {
		log.Error().Err(err).Msg("failed to check bot access")

		return nil, runtime.Error(http.StatusForbidden, "failed to check bot access")
	}

	if err := c.service.BotDeactivate(ctx, req.GetBotId()); err != nil {
		log.Error().Err(err).Msg("failed to deactivate bot")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to deactivate bot")
	}

	return &emptypb.Empty{}, nil
}
