package model

type User struct {
	TelegramModel
	LanguageCode       string        `gorm:"not null"`
	NotificationChatID int64         `gorm:"not null"`
	Transactions       []Transaction `gorm:"foreignKey:UserID;references:TelegramID"`
	Bots               []Bot         `gorm:"foreignKey:UserID;references:TelegramID"`
}
