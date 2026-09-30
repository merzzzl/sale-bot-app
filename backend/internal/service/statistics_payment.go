package service

import (
	"context"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) StatisticsPayment(ctx context.Context, botID int64) error {
	log.Debug().Str("layer", "service").Str("func", "StatisticsPayment").Int64("bot_id", botID).Msg("call service method")

	s.db.WithContext(ctx).Create(&model.Statistic{
		BotID: botID,
		Type:  model.StatisticTypePayments,
		Value: 1,
	})

	return nil
}
