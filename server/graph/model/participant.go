package model

type Participant struct {
	ID      string `json:"id"`
	UserId  int    `json:"userId"`
	EventID int    `json:"eventId"`
	Role    string `json:"role"`
}

type NewParticipant struct {
	UserID  int    `json:"userId"`
	EventID int    `json:"eventId"`
	Role    string `json:"role"`
}
