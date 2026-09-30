package controller

import (
	"context"
	"net/http"
	"net/url"

	"github.com/ltbots/backend/internal/model"
	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
)

func (c *ControllerService) ProductUpdate(ctx context.Context, req *api.ProductUpdateRequest) (*api.Product, error) {
	if err := c.access(ctx, accessCheck{productID: req.GetProductId()}); err != nil {
		log.Error().Err(err).Msg("failed to check product access")

		return nil, runtime.Error(http.StatusForbidden, "failed to check product access")
	}

	if !req.GetUseInvoice() && (req.GetCurrency() == api.Currency_XTR || req.GetPaymentToken() != "") {
		return nil, runtime.Error(http.StatusBadRequest, "use invoice must be true if currency is XTR")
	}

	if !req.GetUseInvoice() {
		_, err := url.Parse(req.GetPayLink())
		if err != nil {
			return nil, runtime.Error(http.StatusBadRequest, "pay link is invalid")
		}
	}

	params := model.UpdateProductParams{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		PayLink:     req.GetPayLink(),
		ImageURL:    req.GetImageUrl(),
		Currency:    req.GetCurrency().String(),
		UseInvoice:  req.GetUseInvoice(),
		PaymentToken: req.GetPaymentToken(),
	}

	if err := c.service.ProductUpdate(ctx, req.GetProductId(), params); err != nil {
		log.Error().Err(err).Msg("failed to update product")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to update product")
	}

	product, err := c.service.ProductGet(ctx, req.GetProductId())
	if err != nil {
		log.Error().Err(err).Msg("failed to get product after update")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to get product after update")
	}

	return &api.Product{
		ProductId:   product.AppID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		PayLink:     product.PayLink,
		ImageUrl:    product.ImageURL,
		BotId:       product.BotID,
		Currency:    api.Currency(api.Currency_value[product.Currency]),
		UseInvoice:  product.UseInvoice,
	}, nil
}
