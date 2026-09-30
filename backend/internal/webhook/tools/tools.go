package tools

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ltbots/backend/internal/service"
	"github.com/sashabaranov/go-openai"
)

type Agent struct {
	service *service.Service
	bot     *bot.Bot
	update  *models.Update
}

func NewAgent(service *service.Service, bot *bot.Bot, update *models.Update) *Agent {
	return &Agent{service: service, bot: bot, update: update}
}

type Tool struct {
	tool    *openai.Tool
	handler func(ctx context.Context, args string) (string, error)
}

func (t *Tool) callHandler(ctx context.Context, call openai.ToolCall) (*openai.ChatCompletionMessage, error) {
	res, err := t.handler(ctx, call.Function.Arguments)
	if err != nil {
		return nil, err
	}

	return &openai.ChatCompletionMessage{
		Role:       openai.ChatMessageRoleTool,
		Content:    res,
		ToolCallID: call.ID,
	}, nil
}

func (a *Agent) tools() []*Tool {
	return []*Tool{
		a.toolProductList(),
		a.toolProductInfo(),
		a.toolProductInvoice(),
	}
}

func (a *Agent) Tools() []openai.Tool {
	tools := make([]openai.Tool, 0, len(a.tools()))

	for _, tool := range a.tools() {
		tools = append(tools, *tool.tool)
	}

	return tools
}

func (a *Agent) HandleTool(ctx context.Context, call openai.ToolCall) (*openai.ChatCompletionMessage, error) {
	for _, tool := range a.tools() {
		if tool.tool.Function.Name == call.Function.Name {
			if message, err := tool.callHandler(ctx, call); err != nil {
				return &openai.ChatCompletionMessage{
					Role:       openai.ChatMessageRoleTool,
					Content:    "{\"success\": false}",
					ToolCallID: call.ID,
				}, err
			} else {
				return message, nil
			}
		}
	}

	return &openai.ChatCompletionMessage{
		Role:       openai.ChatMessageRoleTool,
		Content:    "{\"success\": false}",
		ToolCallID: call.ID,
	}, nil
}
