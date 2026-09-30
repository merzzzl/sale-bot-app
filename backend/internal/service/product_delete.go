package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) ProductDelete(ctx context.Context, productID int64) error {
	log.Debug().Str("layer", "service").Str("func", "ProductDelete").Int64("product_id", productID).Msg("call service method")

	if err := s.db.WithContext(ctx).Delete(&model.Product{}, productID).Error; err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}
