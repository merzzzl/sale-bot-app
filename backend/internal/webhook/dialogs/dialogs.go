package dialogs

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/ltbots/backend/internal/service"
)

type Agent struct {
	service *service.Service
	bot     *bot.Bot
	update  *models.Update
}

func NewAgent(service *service.Service, bot *bot.Bot, update *models.Update) *Agent {
	return &Agent{service: service, bot: bot, update: update}
}
