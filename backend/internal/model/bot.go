package model

type Bot struct {
	TelegramModel
	Token        string       `gorm:"not null"`
	Prompt       string       `gorm:"type:text"`
	UserID       int64        `gorm:"not null;index"`
	User         User         `gorm:"foreignKey:UserID;references:TelegramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Products     []Product    `gorm:"foreignKey:BotID;references:TelegramID"`
	Chats        []Chat       `gorm:"foreignKey:BotID;references:TelegramID"`
	Active       bool         `gorm:"default:false"`
	PresetID     int64        `gorm:"not null;index"`
	PromptPreset PromptPreset `gorm:"foreignKey:PresetID;references:AppID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CreateBotParams struct {
	Token    string
	Prompt   string
	UserID   int64
	PresetID int64
}

type UpdateBotParams struct {
	Prompt   string
	PresetID int64
}
