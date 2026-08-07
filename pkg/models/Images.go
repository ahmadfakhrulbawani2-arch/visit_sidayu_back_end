package models

type Images struct {
	BaseModel
	ImageUrl    string `json:"image_url" gorm:"not null"`
	FileID      string `json:"file_id" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	CustomName  string `json:"custom_name"`
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
	FileID      string `json:"file_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	CustomName  string `json:"custom_name"`
	Description string `json:"description" binding:"required"`
}
