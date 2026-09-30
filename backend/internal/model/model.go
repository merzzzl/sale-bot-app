package model

import (
	"time"

	"gorm.io/gorm"
)

type TelegramModel struct {
	TelegramID int64 `gorm:"column:id;primarykey;autoIncrement:false;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type AppModel struct {
	AppID     int64 `gorm:"column:id;primarykey;autoIncrement:true;unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
