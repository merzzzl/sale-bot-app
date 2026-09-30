package webhook

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ltbots/backend/internal/webhook/dialogs"
	"github.com/ltbots/backend/internal/webhook/payment"
	"github.com/ltbots/backend/internal/webhook/tools"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
)

func (w *Webhook) handler(ctx context.Context, tgBot *bot.Bot, update *models.Update) {
	log.Debug().Str("layer", "webhook").Str("worker", "main").Msg("call webhook method")

	b, err := w.service.BotGet(ctx, tgBot.ID())
	if err != nil {
		log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to get bot")

		return
	}

	if !b.Active {
		log.Info().Int64("bot_id", tgBot.ID()).Msg("bot is not active")

		return
	}

	paymentAgent := payment.NewAgent(w.service, tgBot, update)
	dialogsAgent := dialogs.NewAgent(w.service, tgBot, update)
	toolsAgent := tools.NewAgent(w.service, tgBot, update)

	if ok, err := paymentAgent.PreCheckout(ctx); err != nil || ok {
		if err != nil {
			log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to pre checkout")
		}

		return
	}

	if ok, err := paymentAgent.SuccessfulPayment(ctx); err != nil || ok {
		if err != nil {
			log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to successful payment")
		}

		return
	}

	if update.Message == nil {
		return
	}

	if err := w.service.ChatCreate(ctx, update.Message.Chat.ID, b.TelegramID); err != nil {
		log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to create chat")

		return
	}

	if err := dialogsAgent.MessageSave(ctx, openai.ChatCompletionChoice{
		Message: openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: update.Message.Text,
			Name:    fmt.Sprintf("%s %s (%s)", update.Message.From.FirstName, update.Message.From.LastName, update.Message.From.Username),
		},
	}, ""); err != nil {
		log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to save message")

		return
	}

	history, err := dialogsAgent.HistoryLoad(ctx)
	if err != nil {
		log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to get message list")

		return
	}

	for {
		log.Debug().Int64("bot_id", tgBot.ID()).Int64("chat_id", update.Message.Chat.ID).Msg("creating chat completion")

		response, err := w.openai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
			Model:               w.cahtModel,
			Messages:            history,
			Tools:               toolsAgent.Tools(),
			MaxCompletionTokens: 2048,
		})
		if err != nil {
			log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to create chat completion")

			return
		}

		if err := dialogsAgent.MessageSave(ctx, response.Choices[0], ""); err != nil {
			log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to save message")

			return
		}

		if len(response.Choices[0].Message.ToolCalls) == 0 {
			if _, err := tgBot.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   response.Choices[0].Message.Content,
			}); err != nil {
				log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to send message")

				return
			}

			break
		}

		history = append(history, response.Choices[0].Message)

		for _, toolCall := range response.Choices[0].Message.ToolCalls {
			log.Debug().Int64("bot_id", tgBot.ID()).Int64("chat_id", update.Message.Chat.ID).Str("tool_name", toolCall.Function.Name).Msg("handling tool")

			message, err := toolsAgent.HandleTool(ctx, toolCall)
			if err != nil {
				log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to handle tool")
			}

			history = append(history, *message)

			if err := dialogsAgent.MessageSave(ctx, openai.ChatCompletionChoice{
				Message: *message,
			}, toolCall.ID); err != nil {
				log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to save message")

				return
			}
		}
	}

	if err := w.service.TransactionPayment(ctx, w.messagePrice, b.UserID); err != nil {
		log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to create payment")

		if err := w.service.BotDeactivate(ctx, b.TelegramID); err != nil {
			log.Error().Int64("bot_id", tgBot.ID()).Err(err).Msg("failed to deactivate bot")

			return
		}

		return
	}
}
