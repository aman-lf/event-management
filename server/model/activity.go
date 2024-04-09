package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Name        string
	StartTime   *time.Time
	EndTime     *time.Time
	Description string
	EventID     uint
	Event       Event `gorm:"foreignKey:EventID"`
}
