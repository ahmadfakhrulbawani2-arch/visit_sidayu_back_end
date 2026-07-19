package models

type Superadmins struct {
	BaseModel
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
}

type SuperadminsLoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
