package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ShopsAndUmkmsBlogs struct {
	BaseModel
	Title                 string         `json:"title" gorm:"not null"`
	Slug                  string         `json:"slug" gorm:"unique,not null"`
	Content               string         `json:"content" gorm:"not null"`
	Location              string         `json:"location,omitempty"`
	Rating                float64        `json:"rating,omitempty"`
	Revenue               float64        `json:"revenue,omitempty"`
	MarketedProducts      pq.StringArray `json:"marketed_products" gorm:"type:text[];not null"`
	SalesRatesPiecePerDay int            `json:"sales_rates_per_day,omitempty"`
	ThumbnailID           *uuid.UUID     `json:"thumbnail_id,omitempty"`
	Thumbnail             *Images        `json:"thumbnail,omitempty" gorm:"foreignKey:ThumbnailID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type CreateShopsAndUmkmsBlogsReq struct {
	Title                 string     `json:"title" binding:"required"`
	Content               string     `json:"content" binding:"required"`
	Location              string     `json:"location"`
	Rating                float64    `json:"rating"`
	Revenue               float64    `json:"revenue"`
	MarketedProducts      []string   `json:"marketed_products" binding:"required"`
	SalesRatesPiecePerDay int        `json:"sales_rates_per_day"`
	ThumbnailID           *uuid.UUID `json:"thumbnail_id"`
}
