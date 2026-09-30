package controller

import (
	"context"
	"fmt"
)

type accessCheck struct {
	botID     int64
	productID int64
}

func (c *ControllerService) access(ctx context.Context, check accessCheck) error {
	initData, err := GetInitDataFromContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to get init data: %w", err)
	}

	if check.botID != 0 {
		if err := c.service.BotAccess(ctx, initData.User.ID, check.botID); err != nil {
			return fmt.Errorf("failed to check bot access: %w", err)
		}
	}

	if check.productID != 0 {
		if err := c.service.ProductAccess(ctx, initData.User.ID, check.productID); err != nil {
			return fmt.Errorf("failed to check product access: %w", err)
		}
	}

	return nil
}
