package model

type Expense struct {
	ID          string  `json:"id"`
	ItemName    string  `json:"itemName"`
	Cost        int     `json:"cost"`
	Description *string `json:"description,omitempty"`
	Type        string  `json:"type"`
	EventID     int     `json:"eventID"`
	Event       *Event  `json:"event"`
}

type NewExpense struct {
	ItemName    string  `json:"itemName"`
	Cost        int     `json:"cost"`
	Description *string `json:"description,omitempty"`
	Type        string  `json:"type"`
	EventID     int     `json:"eventID"`
}
