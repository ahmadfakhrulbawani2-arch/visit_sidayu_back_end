package models

import (
	"time"

	"github.com/google/uuid"
)

// TimelinesElement
type TimelinesElement struct {
	BaseModel
	Name             string    `json:"name" gorm:"not null"`
	TimelineDatetime time.Time `json:"timeline_datetime" gorm:"not null"`
	Description      string    `json:"description,omitempty"`
	ExternalLink     string    `json:"external_link,omitempty"`
	TimelinesID      uuid.UUID `json:"timelines_id" gorm:"not null"`
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
	Name          string             `json:"name" gorm:"not null"`
	Description   string             `json:"description,omitempty"`
	BlogID        *uuid.UUID         `json:"blog_id,omitempty" gorm:"uniqueIndex"`
	CultureBlogID *uuid.UUID         `json:"culture_blog_id,omitempty" gorm:"uniqueIndex"`
	Blog          *Blogs             `json:"blog,omitempty" gorm:"foreignKey:BlogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CultureBlog   *CultureBlogs      `json:"culture_blog,omitempty" gorm:"foreignKey:CultureBlogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TimelineData  []TimelinesElement `json:"timeline_data" gorm:"foreignKey:TimelinesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateTimelines struct {
	Name         string                   `json:"name" binding:"required"`
	Description  string                   `json:"description"`
	TimelineData []CreateTimelinesElement `json:"timeline_data" binding:"required,dive"`
}
