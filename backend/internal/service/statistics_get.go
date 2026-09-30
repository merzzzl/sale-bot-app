package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
)

func (s *Service) StatisticsGet(ctx context.Context, botID int64, startTime time.Time, endTime time.Time) ([]model.Statistic, error) {
	log.Debug().Str("layer", "service").Str("func", "StatisticsGet").Int64("bot_id", botID).Time("start_time", startTime).Time("end_time", endTime).Msg("call service method")

	var stats []model.Statistic

	if err := s.db.WithContext(ctx).Where("bot_id = ?", botID).Where("created_at >= ?", startTime).Where("created_at <= ?", endTime).Find(&stats).Error; err != nil {
		return nil, fmt.Errorf("failed to get bot: %w", err)
	}

	return stats, nil
}
