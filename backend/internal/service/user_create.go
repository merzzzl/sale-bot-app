package service

import (
	"context"
	"fmt"

	"github.com/ltbots/backend/internal/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *Service) UserCreate(ctx context.Context, telegramID int64, telegramChatID int64, languageCode string) error {
	log.Debug().Str("layer", "service").Str("func", "UserCreate").Int64("telegram_id", telegramID).Int64("telegram_chat_id", telegramChatID).Str("language_code", languageCode).Msg("call service method")

	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var user model.User

		err := tx.Find(&user, telegramID).Error
		if err != nil {
			return fmt.Errorf("failed to find user: %w", err)
		}

		if (user.NotificationChatID == telegramChatID || telegramChatID == 0) && user.TelegramID == telegramID {
			return nil
		}

		if user.TelegramID == 0 {
			user = model.User{
				LanguageCode: languageCode,
				TelegramModel: model.TelegramModel{
					TelegramID: telegramID,
				},
				NotificationChatID: telegramChatID,
			}

			if err := tx.Create(&user).Error; err != nil {
				return fmt.Errorf("failed to create user: %w", err)
			}
		} else {
			if telegramChatID != 0 {
				if err := tx.Model(&model.User{}).Where("id = ?", telegramID).UpdateColumn("notification_chat_id", telegramChatID).Error; err != nil {
					return fmt.Errorf("failed to update user: %w", err)
				}
			}

			if languageCode != "" {
				if err := tx.Model(&model.User{}).Where("id = ?", telegramID).UpdateColumn("language_code", languageCode).Error; err != nil {
					return fmt.Errorf("failed to update user: %w", err)
				}
			}
		}

		return nil
	}); err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	return nil
}
