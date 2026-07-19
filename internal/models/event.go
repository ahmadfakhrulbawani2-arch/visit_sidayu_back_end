package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	BaseModel
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	UserId      uuid.UUID `json:"user_id" gorm:"not null"`
	User        User      `gorm:"foreignKey:UserId" json:"user"` // Event to User is many to one
	ImageID     uuid.UUID `json:"image_id" gorm:"not null"`
	Image       Image     `json:"image" gorm:"foreignkey:ImageID"` // Event to Image is one to one
	DateTime    time.Time `json:"datetime" binding:"required"`
}

type CreateEventRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	ImageID     uuid.UUID `json:"image_id" binding:"required"` // client will receive new uploaded image url
}
