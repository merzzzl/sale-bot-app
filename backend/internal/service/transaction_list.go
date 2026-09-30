package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) TransactionList(ctx context.Context, userID int64) ([]model.Transaction, error) {
	log.Debug().Str("layer", "service").Str("func", "TransactionList").Int64("user_id", userID).Msg("call service method")

	user := &model.User{}
	if err := s.db.WithContext(ctx).Preload("Transactions").First(user, userID).Error; err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user.Transactions, nil
}
