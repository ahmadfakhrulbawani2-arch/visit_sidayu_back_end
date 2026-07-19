package models

type Image struct {
	BaseModel
	FileID string `json:"file_id" gorm:"unique;not null"`
	Name string `json:"name" gorm:"not null"`
	ImageUrl string `json:"image_url" gorm:"not null"`
}
