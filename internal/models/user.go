package models

type User struct {
	BaseModel
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password,omitempty"`
	Events   []Event `json:"events,omitempty"` // User to Event is one to many
}
