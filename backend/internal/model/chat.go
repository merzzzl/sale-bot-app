package model

type Chat struct {
	TelegramModel
	BotID    int64     `gorm:"primarykey;autoIncrement:false;not null"`
	Bot      Bot       `gorm:"foreignKey:BotID;references:TelegramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Messages []Message `gorm:"foreignKey:ChatID;references:TelegramID"`
}
