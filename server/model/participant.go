package model

import "gorm.io/gorm"

type Participant struct {
	gorm.Model
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
	EventID uint
	Event   Event `gorm:"foreignKey:EventID"`
	Role    string
}
