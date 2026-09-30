package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/sashabaranov/go-openai"
)

func (a *Agent) toolProductInfo() *Tool {
	return &Tool{
		tool: &openai.Tool{
			Type: openai.ToolTypeFunction,
			Function: &openai.FunctionDefinition{
				Name:        "product_info",
				Description: "Информация о продукте",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"product_id": map[string]any{
							"type":        "integer",
							"description": "ID продукта, о котором требуется информация",
						},
					},
				},
			},
		},
		handler: a.toolProductInfoHandler,
	}
}

func (a *Agent) toolProductInfoHandler(ctx context.Context, args string) (string, error) {
	arguments := struct {
		ProductID int64 `json:"product_id"`
	}{}

	if err := json.Unmarshal([]byte(args), &arguments); err != nil {
		return "", fmt.Errorf("failed to unmarshal arguments: %w", err)
	}

	products, err := a.service.ProductList(ctx, a.bot.ID())
	if err != nil {
		return "", fmt.Errorf("failed to get product list: %w", err)
	}

	var product *model.Product

	for _, p := range products {
		if p.AppID == arguments.ProductID {
			product = &p

			break
		}
	}

	if product == nil {
		return "{\"success\": false}", nil
	}

	productInfo := struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       int64  `json:"price"`
		Currency    string `json:"currency"`
	}{
		ID:          product.AppID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price / 100,
		Currency:    product.Currency,
	}

	js, err := json.Marshal(productInfo)
	if err != nil {
		return "", fmt.Errorf("failed to marshal product list: %w", err)
	}

	return string(js), nil
}
