package models

type Roles struct {
	BaseModel
	Name        string `json:"name" gorm:"unique;not null"`
	Level       int    `json:"level" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}

type CreateRoles struct {
	Name        string `json:"name" binding:"required"`
	Level       int    `json:"level" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateRoles struct {
	Name        *string `json:"name"`
	Level       *int    `json:"level"`
	Description *string `json:"description"`
}
