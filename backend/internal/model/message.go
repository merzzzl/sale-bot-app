package model

type MessageType string

const (
	MessageTypeText    MessageType = "text"
	MessageTypeSummary MessageType = "summary"
)

type MessageSenderType string

const (
	MessageSenderTypeUser      MessageSenderType = "user"
	MessageSenderTypeAssistant MessageSenderType = "assistant"
	MessageSenderTypeSystem    MessageSenderType = "system"
	MessageSenderTypeTool      MessageSenderType = "tool"
)

type Message struct {
	AppModel
	ChatID     int64             `gorm:"not null;index"`
	Chat       Chat              `gorm:"foreignKey:ChatID;references:TelegramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SenderName string            `gorm:"not null"`
	SenderType MessageSenderType `gorm:"not null"`
	Type       MessageType       `gorm:"default:text"`
	Text       string            `gorm:"type:text"`
	ToolCallID string            `gorm:"default:''"`
	ToolCalls  []ToolCall        `gorm:"serializer:json"`
}

type CreateMessageParams struct {
	BotID      int64
	ChatID     int64
	SenderName string
	SenderType MessageSenderType
	Type       MessageType
	Text       string
	ToolCallID string
	ToolCalls  []ToolCall
}

type ToolCall struct {
	ID       string
	ToolName string
	ToolArgs string
}
