package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) ProductAccess(ctx context.Context, userID, productID int64) error {
	log.Debug().Str("layer", "service").Str("func", "ProductAccess").Int64("user_id", userID).Int64("product_id", productID).Msg("call service method")

	if err := s.db.WithContext(ctx).Preload("Bot", "user_id = ?", userID).First(&model.Product{}, productID).Error; err != nil {
		return fmt.Errorf("failed to get product: %w", err)
	}

	return nil
}
