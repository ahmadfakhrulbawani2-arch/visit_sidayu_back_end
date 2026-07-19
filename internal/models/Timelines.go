package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TimelinesElement
type TimelinesElement struct {
	ID               uuid.UUID `json:"ID" gorm:"primaryKey"`
	Name             string    `json:"name" gorm:"not null"`
	TimelineDatetime time.Time `json:"timeline_datetime" gorm:"not null"`
	Description      string    `json:"description,omitempty"`
	ExternalLink     string    `json:"external_link,omitempty"`
	TimelinesID      uuid.UUID `json:"timelines_id" gorm:"not null"`
}

func (t *TimelinesElement) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}

type CreateTimelinesElement struct {
	Name             string    `json:"name" binding:"required"`
	TimelineDatetime time.Time `json:"timeline_datetime" binding:"required"`
	Description      string    `json:"description"`
	ExternalLink     string    `json:"external_link"`
}

// Timelines
type Timelines struct {
	BaseModel
	Name         string             `json:"name" gorm:"not null"`
	Description  string             `json:"description,omitempty"`
	BlogID       uuid.UUID          `json:"blog_id" gorm:"not null;uniqueIndex"`
	Blog         *Blogs             `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TimelineData []TimelinesElement `json:"timeline_data" gorm:"foreignKey:TimelinesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateTimelines struct {
	Name         string                   `json:"name" binding:"required"`
	Description  string                   `json:"description"`
	TimelineData []CreateTimelinesElement `json:"timeline_data" binding:"required,dive"`
}
