package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) TransactionAmount(ctx context.Context, userID int64) (int64, error) {
	log.Debug().Str("layer", "service").Str("func", "TransactionAmount").Int64("user_id", userID).Msg("call service method")

	user := &model.User{}
	if err := s.db.WithContext(ctx).Preload("Transactions").First(user, userID).Error; err != nil {
		return 0, fmt.Errorf("failed to get user: %w", err)
	}

	var totalAmount int64
	for _, transaction := range user.Transactions {
		totalAmount += transaction.Amount
	}

	return totalAmount, nil
}
