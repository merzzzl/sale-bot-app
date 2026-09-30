package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *Service) TransactionPayment(ctx context.Context, amount int64, userID int64) error {
	log.Debug().Str("layer", "service").Str("func", "TransactionPayment").Int64("amount", amount).Int64("user_id", userID).Msg("call service method")

	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		user := &model.User{}
		if err := tx.Preload("Transactions").First(user, userID).Error; err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}

		var totalAmount int64
		for _, transaction := range user.Transactions {
			totalAmount += transaction.Amount
		}

		if totalAmount < amount {
			return fmt.Errorf("insufficient balance")
		}

		transaction := &model.Transaction{
			UserID: userID,
			Amount: -amount,
			Type:   model.TransactionTypePayment,
		}

		if err := tx.Create(transaction).Error; err != nil {
			return fmt.Errorf("failed to create payment: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	return nil
}
