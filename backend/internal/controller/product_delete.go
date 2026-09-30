package controller

import (
	"context"
	"net/http"

	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *ControllerService) ProductDelete(ctx context.Context, req *api.ProductDeleteRequest) (*emptypb.Empty, error) {
	if err := c.access(ctx, accessCheck{productID: req.GetProductId()}); err != nil {
		log.Error().Err(err).Msg("failed to check product access")

		return nil, runtime.Error(http.StatusForbidden, "failed to check product access")
	}

	if err := c.service.ProductDelete(ctx, req.GetProductId()); err != nil {
		log.Error().Err(err).Msg("failed to delete product")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to delete product")
	}

	return &emptypb.Empty{}, nil
}
