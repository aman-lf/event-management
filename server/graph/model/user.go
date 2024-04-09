package model

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	PhoneNo string `json:"phoneNo"`
}

type NewUser struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	PhoneNo *string `json:"phoneNo,omitempty"`
}
