package model

type StatisticType string

const (
	StatisticTypeMessages StatisticType = "messages"
	StatisticTypeChats    StatisticType = "chats"
	StatisticTypePayments StatisticType = "payments"
)

type Statistic struct {
	AppModel
	BotID int64         `gorm:"not null;index"`
	Bot   Bot           `gorm:"foreignKey:BotID;references:TelegramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Type  StatisticType `gorm:"not null"`
	Value int64         `gorm:"not null;default:0"`
}
