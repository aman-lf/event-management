package model

import (
	"github.com/aman-lf/event-management/data"
	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
	EventID uint
	Event   Event `gorm:"foreignKey:EventID"`
	Role    string
}

func (p *Participant) IsAdmin(eventId int) bool {
	return p.Role == data.ADMIN
}

func (p *Participant) IsContributor(eventId int) bool {
	return p.Role == data.CONTRIBUTOR
}

func (p *Participant) IsAttendee(eventId int) bool {
	return p.Role == data.ATTENDEE
}
