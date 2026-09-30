package model

type TransactionType string

const (
	TransactionTypeDeposit TransactionType = "deposit"
	TransactionTypePayment TransactionType = "payment"
)

type Transaction struct {
	AppModel
	UserID int64           `gorm:"not null;index"`
	User   User            `gorm:"foreignKey:UserID;references:TelegramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Amount int64           `gorm:"not null"`
	Type   TransactionType `gorm:"not null"`
}
