package service

import (
	"github.com/go-telegram/bot"
	"gorm.io/gorm"
)

type Service struct {
	db        *gorm.DB
	mainBot   *bot.Bot
	botEvents chan any
}

func NewService(db *gorm.DB, mainBot *bot.Bot) *Service {
	s := &Service{
		db:        db,
		mainBot:   mainBot,
		botEvents: make(chan any, 10),
	}

	s.botEvents <- struct{}{}

	return s
}

func (s *Service) BotEvents() <-chan any {
	return s.botEvents
}
