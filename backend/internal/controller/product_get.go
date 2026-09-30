package controller

import (
	"context"
	"net/http"

	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
)

func (c *ControllerService) ProductGet(ctx context.Context, req *api.ProductGetRequest) (*api.Product, error) {
	if err := c.access(ctx, accessCheck{productID: req.GetProductId()}); err != nil {
		log.Error().Err(err).Msg("failed to check product access")

		return nil, runtime.Error(http.StatusForbidden, "failed to check product access")
	}

	product, err := c.service.ProductGet(ctx, req.GetProductId())
	if err != nil {
		log.Error().Err(err).Msg("failed to get product")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to get product")
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
