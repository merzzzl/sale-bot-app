package controller

import (
	"context"
	"net/http"

	"github.com/ltbots/protocols/go/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *ControllerService) PromptPresetList(ctx context.Context, req *emptypb.Empty) (*api.PromptPresetListResponse, error) {
	presets, err := c.service.PromptPresetList(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get preset list")

		return nil, runtime.Error(http.StatusInternalServerError, "failed to get preset list")
	}

	var response api.PromptPresetListResponse

	for _, preset := range presets {
		response.PromptPresets = append(response.PromptPresets, &api.PromptPreset{
			PresetId:    preset.AppID,
			Name:        preset.Name,
			Description: preset.Description,
		})
	}

	return &response, nil
}
