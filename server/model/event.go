package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string
	StartDate   *time.Time
	EndDate     *time.Time
	Location    string
	Type        string
	Description string
}
