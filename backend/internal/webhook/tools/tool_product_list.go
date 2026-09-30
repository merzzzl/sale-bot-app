package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func (a *Agent) toolProductList() *Tool {
	return &Tool{
		tool: &openai.Tool{
			Type: openai.ToolTypeFunction,
			Function: &openai.FunctionDefinition{
				Name:        "product_list",
				Description: "Список доступных товаров/услуг",
				Parameters: map[string]any{
					"type":       "object",
					"properties": map[string]any{},
				},
			},
		},
		handler: a.toolProductListHandler,
	}
}

func (a *Agent) toolProductListHandler(ctx context.Context, args string) (string, error) {
	products, err := a.service.ProductList(ctx, a.bot.ID())
	if err != nil {
		return "", fmt.Errorf("failed to get product list: %w", err)
	}

	type productList struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	tinyProducts := []productList{}

	for _, product := range products {
		tinyProducts = append(tinyProducts, productList{
			ID:   product.AppID,
			Name: product.Name,
		})
	}

	js, err := json.Marshal(tinyProducts)
	if err != nil {
		return "", fmt.Errorf("failed to marshal product list: %w", err)
	}

	return string(js), nil
}
