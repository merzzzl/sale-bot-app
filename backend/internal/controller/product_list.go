package controller

import (
	"context"
	"net/http"

	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
)

func (c *ControllerService) ProductList(ctx context.Context, req *api.ProductListRequest) (*api.ProductListResponse, error) {
	if err := c.access(ctx, accessCheck{botID: req.GetBotId()}); err != nil {
		log.Error().Err(err).Msg("failed to check bot access")

		return nil, runtime.Error(http.StatusForbidden, "failed to check bot access")
	}

	products, err := c.service.ProductList(ctx, req.GetBotId())
	if err != nil {
		log.Error().Err(err).Msg("failed to get product list")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to get product list")
	}

	var response api.ProductListResponse

	for _, product := range products {
		response.Products = append(response.Products, &api.Product{
			ProductId:   product.AppID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			PayLink:     product.PayLink,
			ImageUrl:    product.ImageURL,
			BotId:       product.BotID,
			Currency:    api.Currency(api.Currency_value[product.Currency]),
			UseInvoice:  product.UseInvoice,
		})
	}

	return &response, nil
}
