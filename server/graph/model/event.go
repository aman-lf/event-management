package model

type Event struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	StartDate   string  `json:"startDate"`
	EndDate     string  `json:"endDate"`
	Location    string  `json:"location"`
	Type        *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
}

type NewEvent struct {
	Name        string  `json:"name"`
	StartDate   string  `json:"startDate"`
	EndDate     string  `json:"endDate"`
	Location    string  `json:"location"`
	Type        *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
}
