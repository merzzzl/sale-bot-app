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

func (c *ControllerService) ProductCreate(ctx context.Context, req *api.ProductCreateRequest) (*api.Product, error) {
	if err := c.access(ctx, accessCheck{botID: req.GetBotId()}); err != nil {
		log.Error().Err(err).Msg("failed to check bot access")

		return nil, runtime.Error(http.StatusForbidden, "failed to check bot access")
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

	params := model.CreateProductParams{
		BotID:       req.GetBotId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		PayLink:     req.GetPayLink(),
		ImageURL:    req.GetImageUrl(),
		Currency:    req.GetCurrency().String(),
		UseInvoice:  req.GetUseInvoice(),
		PaymentToken: req.GetPaymentToken(),
	}

	product, err := c.service.ProductCreate(ctx, params)
	if err != nil {
		log.Error().Err(err).Msg("failed to create product")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to create product")
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
