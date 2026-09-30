package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) UserGet(ctx context.Context, userID int64) (*model.User, error) {
	log.Debug().Str("layer", "service").Str("func", "UserGet").Int64("user_id", userID).Msg("call service method")

	var user model.User

	if err := s.db.WithContext(ctx).Find(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}
