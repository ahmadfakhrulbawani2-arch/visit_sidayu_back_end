package models

import "github.com/google/uuid"

type Galleries struct {
	BaseModel
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Slug        string    `json:"slug" gorm:"unique;not null"` // in FE it's not mandatory, BE will auto fill if no slug provided
	ImageID     uuid.UUID `json:"image_id" gorm:"not null"`
	Image       Images    `json:"image" gorm:"foreignKey:ImageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type CreateGalleries struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Slug        string    `json:"slug"`
	ImageID     uuid.UUID `json:"image_id" binding:"required"`
}
