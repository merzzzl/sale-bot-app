package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ltbots/backend/internal/i18n"
	"github.com/ltbots/backend/internal/model"
	"github.com/sashabaranov/go-openai"
)

func (a *Agent) toolProductInvoice() *Tool {
	return &Tool{
		tool: &openai.Tool{
			Type: openai.ToolTypeFunction,
			Function: &openai.FunctionDefinition{
				Name:        "product_invoice",
				Description: "Отправить форму оплаты для продукта",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"product_id": map[string]any{
							"type":        "integer",
							"description": "ID продукта, для которого требуется отправить форму оплаты",
						},
					},
				},
			},
		},
		handler: a.toolProductInvoiceHandler,
	}
}

func (a *Agent) toolProductInvoiceHandler(ctx context.Context, args string) (string, error) {
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

	payload := struct {
		ProductID int64  `json:"product_id"`
		Timestamp string `json:"timestamp"`
	}{
		ProductID: product.AppID,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	js, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	if product.Currency == "XTR" {
		product.Price /= 100
	}

	if product.UseInvoice {
		if _, err := a.bot.SendInvoice(ctx, &bot.SendInvoiceParams{
			ChatID:        a.update.Message.Chat.ID,
			Title:         product.Name,
			Description:   product.Description,
			Payload:       string(js),
			Currency:      product.Currency,
			ProviderToken: product.PaymentToken,
			Prices: []models.LabeledPrice{
				{
					Label:  product.Name,
					Amount: int(product.Price),
				},
			},
		}); err != nil {
			return "", fmt.Errorf("failed to send invoice: %w", err)
		}
	} else {
		if _, err := a.bot.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: a.update.Message.Chat.ID,
			Text:   fmt.Sprintf("%s\n\n%s", product.Name, product.Description),
			Entities: []models.MessageEntity{
				{
					Type:   models.MessageEntityTypeBold,
					Offset: 0,
					Length: utf8.RuneCountInString(product.Name),
				},
				{
					Type:   models.MessageEntityTypeItalic,
					Offset: utf8.RuneCountInString(product.Name) + 2,
					Length: utf8.RuneCountInString(product.Description),
				},
			},
			ReplyMarkup: &models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{
					{
						{
							Text:         fmt.Sprintf(i18n.Localize(a.update.Message.From.LanguageCode, "handler_pay_button"), product.Price/100, product.Currency),
							URL:          product.PayLink,
							CallbackData: string(js),
						},
					},
				},
			},
		}); err != nil {
			return "", fmt.Errorf("failed to send message: %w", err)
		}
	}

	return "{\"success\": true}", nil
}
