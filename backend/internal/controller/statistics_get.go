package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/ltbots/backend/internal/model"
	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
)

func (c *ControllerService) StatisticsGet(ctx context.Context, req *api.StatisticsGetRequest) (*api.StatisticsGetResponse, error) {
	if err := c.access(ctx, accessCheck{botID: req.GetBotId()}); err != nil {
		log.Error().Err(err).Msg("failed to check bot access")

		return nil, runtime.Error(http.StatusForbidden, "failed to check bot access")
	}

	startTime := time.Unix(req.GetStartTime(), 0)
	endTime := time.Unix(req.GetEndTime(), 0)

	if endTime.Before(startTime) {
		log.Error().Msg("end time must be after start time")

		return nil, runtime.Error(http.StatusBadRequest, "end time must be after start time")
	}

	if endTime.Sub(startTime) > time.Hour*24*90 {
		log.Error().Msg("time range must be less than 30 days")

		return nil, runtime.Error(http.StatusBadRequest, "time range must be less than 30 days")
	}

	stats, err := c.service.StatisticsGet(ctx, req.GetBotId(), startTime, endTime)
	if err != nil {
		log.Error().Err(err).Msg("failed to get statistics")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to get statistics")
	}

	var response api.StatisticsGetResponse

	for _, stat := range stats {
		var statType api.StatisticType

		switch stat.Type {
		case model.StatisticTypeMessages:
			statType = api.StatisticType_MESSAGES
		case model.StatisticTypeChats:
			statType = api.StatisticType_CHATS
		case model.StatisticTypePayments:
			statType = api.StatisticType_PAYMENTS
		}
		response.Records = append(response.Records, &api.StatisticsGetResponse_Record{
			Type:      statType,
			Value:     stat.Value,
			Timestamp: stat.CreatedAt.Unix(),
		})
	}

	return &response, nil
}
