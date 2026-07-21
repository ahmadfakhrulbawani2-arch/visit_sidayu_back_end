package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type CultureBlogs struct {
	BaseModel
	Title                    string         `json:"title" gorm:"not null"`
	Description              string         `json:"description,omitempty"`
	Content                  string         `json:"content" gorm:"not null"`
	Tags                     pq.StringArray `json:"tags" gorm:"type:text[];not null"`
	ThemeType                string         `json:"theme_type" gorm:"not null"`
	ThumbnailID              *uuid.UUID     `json:"thumbnail_id,omitempty"`
	Thumbnail                *Images        `json:"thumbnail,omitempty" gorm:"foreignKey:ThumbnailID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Location                 string         `json:"location,omitempty"`
	Author                   string         `json:"author" gorm:"not null"`
	BlogWrittenDatetime      time.Time      `json:"blog_written_datetime" gorm:"not null"`
	EstimatedMinutesReadTime int            `json:"estimated_read_time" gorm:"not null"`
	ExternalLinks            pq.StringArray `json:"external_links,omitempty" gorm:"type:text[]"`
	Timeline                 *Timelines     `json:"timeline,omitempty" gorm:"foreignKey:CultureBlogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateCultureBlogsReq struct {
	Title                    string           `json:"title" binding:"required"`
	Description              string           `json:"description"`
	Content                  string           `json:"content" binding:"required"`
	Tags                     []string         `json:"tags" binding:"required"`
	ThemeType                string           `json:"theme_type" binding:"required"`
	ThumbnailID              *uuid.UUID       `json:"thumbnail_id"`
	Location                 string           `json:"location"`
	Author                   string           `json:"author" binding:"required"`
	BlogWrittenDatetime      time.Time        `json:"blog_written_datetime" binding:"required"`
	EstimatedMinutesReadTime int              `json:"estimated_read_time"`
	ExternalLinks            []string         `json:"external_links"`
	Timeline                 *CreateTimelines `json:"timeline" binding:"omitempty,dive"`
}
