package dialogs

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/i18n"
	"github.com/ltbots/backend/internal/model"
	"github.com/sashabaranov/go-openai"
)

func (a *Agent) HistoryLoad(ctx context.Context) ([]openai.ChatCompletionMessage, error) {
	messages, err := a.service.MessageList(ctx, a.update.Message.Chat.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get message list: %w", err)
	}

	presets, err := a.service.PromptPresetList(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get prompt preset list: %w", err)
	}

	bot, err := a.service.BotGet(ctx, a.bot.ID())
	if err != nil {
		return nil, fmt.Errorf("failed to get bot: %w", err)
	}

	var systemPrompt string

	for _, preset := range presets {
		if bot.PresetID == preset.AppID {
			systemPrompt = preset.Prompt + "\n\n" + bot.Prompt

			break
		}
	}

	systemPrompt = systemPrompt + "\n\n" + i18n.Localize(a.update.Message.From.LanguageCode, "prompt_language_message")

	var userMessages int
	var aiMessages []openai.ChatCompletionMessage

	aiMessages = append(aiMessages, openai.ChatCompletionMessage{
		Role:    "system",
		Content: systemPrompt,
	})

	for _, message := range messages {
		if userMessages >= 10 {
			break
		}

		if message.SenderType == model.MessageSenderTypeUser {
			userMessages++
		}

		var role string

		switch message.SenderType {
		case model.MessageSenderTypeUser:
			role = openai.ChatMessageRoleUser
		case model.MessageSenderTypeAssistant:
			role = openai.ChatMessageRoleAssistant
		case model.MessageSenderTypeSystem:
			role = openai.ChatMessageRoleSystem
		case model.MessageSenderTypeTool:
			role = openai.ChatMessageRoleTool
		}

		var toolCalls []openai.ToolCall

		for _, toolCall := range message.ToolCalls {
			toolCalls = append(toolCalls, openai.ToolCall{
				ID:   toolCall.ID,
				Type: openai.ToolTypeFunction,
				Function: openai.FunctionCall{
					Name:      toolCall.ToolName,
					Arguments: toolCall.ToolArgs,
				},
			})
		}

		aiMessages = append(aiMessages, openai.ChatCompletionMessage{
			Role:       role,
			Content:    message.Text,
			ToolCalls:  toolCalls,
			ToolCallID: message.ToolCallID,
		})
	}

	return aiMessages, nil
}
