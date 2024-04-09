package model

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	ItemName    string
	Cost        int
	Description *string
	Type        string
	EventID     uint
	Event       Event `gorm:"foreignKey:EventID"`
}
