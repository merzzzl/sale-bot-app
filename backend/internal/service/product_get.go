package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) ProductGet(ctx context.Context, productID int64) (*model.Product, error) {
	log.Debug().Str("layer", "service").Str("func", "ProductGet").Int64("product_id", productID).Msg("call service method")

	var product *model.Product

	if err := s.db.WithContext(ctx).First(&product, productID).Error; err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return product, nil
}
