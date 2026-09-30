package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-telegram/bot"
)

func (a *Agent) PreCheckout(ctx context.Context) (bool, error) {
	if a.update.PreCheckoutQuery == nil {
		return false, nil
	}

	payload := struct {
		ProductID int64  `json:"product_id"`
		Timestamp string `json:"timestamp"`
	}{}

	if err := json.Unmarshal([]byte(a.update.PreCheckoutQuery.InvoicePayload), &payload); err != nil {
		return false, fmt.Errorf("failed to unmarshal invoice payload: %w", err)
	}

	payloadTimestamp, err := time.Parse(time.RFC3339, payload.Timestamp)
	if err != nil {
		return false, fmt.Errorf("failed to parse invoice payload: %w", err)
	}

	if time.Now().Add(24 * time.Hour).After(payloadTimestamp) {
		return false, fmt.Errorf("invoice payload is expired")
	}

	if _, err := a.bot.AnswerPreCheckoutQuery(ctx, &bot.AnswerPreCheckoutQueryParams{
		PreCheckoutQueryID: a.update.PreCheckoutQuery.ID,
		OK:                 true,
	}); err != nil {
		return false, fmt.Errorf("failed to answer pre checkout query: %w", err)
	}

	return true, nil
}
