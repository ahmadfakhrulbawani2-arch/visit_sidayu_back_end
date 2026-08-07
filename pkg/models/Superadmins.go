package models

type Superadmins struct {
	BaseModel
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
}

type SuperadminsLoginPayload struct {
	Identity string `json:"identity" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateSuperadmins struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
}

type UpdateSuperadminDTO struct {
	Username    *string `json:"username"`
	Password    string  `json:"password" binding:"required"`
	NewPassword *string `json:"new_password"`
}
