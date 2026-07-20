package models

type Images struct {
	BaseModel
	ImageUrl    string `json:"image_url" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
}

type GetImages struct {
	BaseModel
	ImageUrl    string `json:"image_url"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type CreateImages struct {
	ImageUrl    string `json:"image_url" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
