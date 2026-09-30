package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *Service) ProductUpdate(ctx context.Context, productID int64, params model.UpdateProductParams) error {
	log.Debug().Str("layer", "service").Str("func", "ProductUpdate").Int64("product_id", productID).Msg("call service method")

	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var product model.Product

		if err := tx.First(&product, productID).Error; err != nil {
			return fmt.Errorf("failed to get product: %w", err)
		}

		updateData := make(map[string]interface{})

		if params.Name != "" && params.Name != product.Name {
			updateData["name"] = params.Name
		}

		if params.Description != "" && params.Description != product.Description {
			updateData["description"] = params.Description
		}

		if params.Price != 0 && params.Price != product.Price {
			updateData["price"] = params.Price
		}

		if params.PayLink != "" && params.PayLink != product.PayLink {
			updateData["pay_link"] = params.PayLink
		}

		if params.ImageURL != "" && params.ImageURL != product.ImageURL {
			updateData["image_url"] = params.ImageURL
		}

		if params.Currency != "" && params.Currency != product.Currency {
			updateData["currency"] = params.Currency
		}

		if params.UseInvoice != product.UseInvoice {
			updateData["use_invoice"] = params.UseInvoice
		}

		if params.PaymentToken != product.PaymentToken {
			updateData["payment_token"] = params.PaymentToken
		}

		if err := tx.Model(&product).Where("id = ?", productID).Updates(updateData).Error; err != nil {
			return fmt.Errorf("failed to update product: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	return nil
}
