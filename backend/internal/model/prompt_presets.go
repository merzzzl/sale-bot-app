package model

type PromptPreset struct {
	AppModel
	Name        string `gorm:"not null;size:255"`
	Description string `gorm:"type:text"`
	Prompt      string `gorm:"type:text"`
}
