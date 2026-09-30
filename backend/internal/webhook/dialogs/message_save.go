package dialogs

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/sashabaranov/go-openai"
)

func (a *Agent) MessageSave(ctx context.Context, choice openai.ChatCompletionChoice, toolCallID string) error {
	var senderType model.MessageSenderType

	switch choice.Message.Role {
	case openai.ChatMessageRoleUser:
		senderType = model.MessageSenderTypeUser
	case openai.ChatMessageRoleAssistant:
		senderType = model.MessageSenderTypeAssistant
	case openai.ChatMessageRoleSystem:
		senderType = model.MessageSenderTypeSystem
	case openai.ChatMessageRoleTool:
		senderType = model.MessageSenderTypeTool
	}

	var toolCalls []model.ToolCall

	for _, toolCall := range choice.Message.ToolCalls {
		toolCalls = append(toolCalls, model.ToolCall{
			ID:       toolCall.ID,
			ToolName: toolCall.Function.Name,
			ToolArgs: toolCall.Function.Arguments,
		})
	}

	var senderName string

	switch senderType {
	case model.MessageSenderTypeUser:
		senderName = fmt.Sprintf("%s %s (%s)", a.update.Message.From.FirstName, a.update.Message.From.LastName, a.update.Message.From.Username)
	case model.MessageSenderTypeAssistant:
		senderName = "assistant"
	case model.MessageSenderTypeSystem:
		senderName = "system"
	case model.MessageSenderTypeTool:
		senderName = "tool"
	}

	if senderType != model.MessageSenderTypeTool && toolCallID != "" {
		return fmt.Errorf("tool call id is empty for non-tool message")
	}

	if err := a.service.MessageCreate(ctx, model.CreateMessageParams{
		BotID:      a.bot.ID(),
		ChatID:     a.update.Message.Chat.ID,
		SenderName: senderName,
		SenderType: senderType,
		Text:       choice.Message.Content,
		ToolCalls:  toolCalls,
		ToolCallID: toolCallID,
	}); err != nil {
		return fmt.Errorf("failed to create message: %w", err)
	}

	return nil
}
