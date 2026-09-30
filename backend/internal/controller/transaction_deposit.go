package controller

import (
	"context"
	"net/http"

	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *ControllerService) TransactionBill(ctx context.Context, req *api.TransactionBillRequest) (*emptypb.Empty, error) {
	initData, err := GetInitDataFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get init data")

		return nil, runtime.Error(http.StatusUnauthorized, "failed to get init data")
	}

	if err := c.service.TransactionBill(ctx, initData.User.ID, req.GetAmount()); err != nil {
		log.Error().Err(err).Msg("failed to bill")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to bill")
	}

	return &emptypb.Empty{}, nil
}
