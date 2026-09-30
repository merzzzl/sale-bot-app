package controller

import (
	"context"
	"net/http"

	"github.com/ltbots/backend/internal/model"
	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *ControllerService) TransactionsList(ctx context.Context, req *emptypb.Empty) (*api.TransactionsListResponse, error) {
	initData, err := GetInitDataFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get init data")

		return nil, runtime.Error(http.StatusUnauthorized, "failed to get init data")
	}

	transactions, err := c.service.TransactionList(ctx, initData.User.ID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get transactions list")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to get transactions list")
	}

	var response api.TransactionsListResponse

	for _, transaction := range transactions {
		var transactionType api.TransactionType
		if transaction.Type == model.TransactionTypeDeposit {
			transactionType = api.TransactionType_DEPOSIT
		} else {
			transactionType = api.TransactionType_PAYMENT
		}

		response.Transactions = append(response.Transactions, &api.Transaction{
			TransactionId: transaction.AppID,
			Amount:        transaction.Amount,
			Type:          transactionType,
			Timestamp:     transaction.CreatedAt.Unix(),
		})
	}

	return &response, nil
}
