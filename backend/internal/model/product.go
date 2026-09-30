package model

type Product struct {
	AppModel
	Name         string `gorm:"not null;size:255"`
	Description  string `gorm:"type:text"`
	Price        int64  `gorm:"not null;check:price >= 0"`
	PayLink      string `gorm:"type:text"`
	ImageURL     string `gorm:"type:text"`
	BotID        int64  `gorm:"not null;index"`
	Bot          Bot    `gorm:"foreignKey:BotID;references:TelegramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Currency     string `gorm:"not null"`
	UseInvoice   bool   `gorm:"default:false"`
	PaymentToken string `gorm:"type:text"`
}

type CreateProductParams struct {
	Name         string
	Description  string
	Price        int64
	PayLink      string
	ImageURL     string
	BotID        int64
	Currency     string
	UseInvoice   bool
	PaymentToken string
}

type UpdateProductParams struct {
	Name         string
	Description  string
	Price        int64
	PayLink      string
	ImageURL     string
	Currency     string
	UseInvoice   bool
	PaymentToken string
}
