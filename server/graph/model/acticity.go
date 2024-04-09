package model

type Activity struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	StartTime   string  `json:"startTime"`
	EndTime     string  `json:"endTime"`
	Description *string `json:"description,omitempty"`
	EventID     int     `json:"eventID"`
	Event       *Event  `json:"event"`
}

type NewActivity struct {
	Name        string  `json:"name"`
	StartTime   string  `json:"startTime"`
	EndTime     string  `json:"endTime"`
	Description *string `json:"description,omitempty"`
	EventID     int     `json:"eventId"`
}
