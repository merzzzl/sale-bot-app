package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) ProductCreate(ctx context.Context, params model.CreateProductParams) (*model.Product, error) {
	log.Debug().Str("layer", "service").Str("func", "ProductCreate").Int64("bot_id", params.BotID).Msg("call service method")

	product := &model.Product{
		Name:         params.Name,
		Description:  params.Description,
		Price:        params.Price,
		PayLink:      params.PayLink,
		ImageURL:     params.ImageURL,
		BotID:        params.BotID,
		Currency:     params.Currency,
		UseInvoice:   params.UseInvoice,
		PaymentToken: params.PaymentToken,
	}

	if err := s.db.WithContext(ctx).Create(product).Error; err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return product, nil
}
