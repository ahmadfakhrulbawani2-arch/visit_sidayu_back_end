package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Blogs struct {
	BaseModel
	Title                string         `json:"title" gorm:"not null"`
	Description          string         `json:"description,omitempty"`
	Tags                 pq.StringArray `json:"tags" gorm:"type:text[];not null"`
	Content              string         `json:"content" gorm:"not null"`
	Author               string         `json:"author" gorm:"not null"`
	BlogWrittenDatetime  time.Time      `json:"blog_written_datetime" gorm:"not null"`
	EstimatedMinutesRead int            `json:"estimated_minutes_read" gorm:"not null"`
	ThumbnailID          *uuid.UUID     `json:"thumbnail_id,omitempty"`
	Thumbnail            *Images        `json:"thumbnail,omitempty" gorm:"foreignKey:ThumbnailID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Location             string         `json:"location,omitempty"`
	ExternalLinks        pq.StringArray `json:"external_links,omitempty" gorm:"type:text[]"`
	Timeline             *Timelines     `json:"timeline,omitempty" gorm:"foreignKey:BlogID;references:ID;"`
}

type CreateBlogs struct {
	Title                string           `json:"title" binding:"required"`
	Description          string           `json:"description"`
	Tags                 []string         `json:"tags" binding:"required"`
	Content              string           `json:"content" binding:"required"`
	Author               string           `json:"author" binding:"required"`
	BlogWrittenDatetime  time.Time        `json:"blog_written_datetime" binding:"required"`
	EstimatedMinutesRead int              `json:"estimated_minutes_read" binding:"required"`
	ThumbnailID          *uuid.UUID       `json:"thumbnail_id"`
	Location             string           `json:"location"`
	ExternalLinks        []string         `json:"external_links"`
	Timeline             *CreateTimelines `json:"timeline"`
}
