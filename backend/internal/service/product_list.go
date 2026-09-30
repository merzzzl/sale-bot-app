package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) ProductList(ctx context.Context, botID int64) ([]model.Product, error) {
	log.Debug().Str("layer", "service").Str("func", "ProductList").Int64("bot_id", botID).Msg("call service method")

	var products []model.Product

	if err := s.db.WithContext(ctx).Where("bot_id = ?", botID).Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return products, nil
}
