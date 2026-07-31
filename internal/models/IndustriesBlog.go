package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type IndustriesBlogs struct {
	BaseModel
	Title                      string         `json:"title" gorm:"not null"`
	Slug                       string         `json:"slug" gorm:"unique,not null"`
	Content                    string         `json:"content" gorm:"not null"`
	Location                   string         `json:"location,omitempty"`
	Rating                     float64        `json:"rating,omitempty"`
	Revenue                    float64        `json:"revenue,omitempty"`
	ProducedProducts           pq.StringArray `json:"produced_products" gorm:"type:text[];not null"`
	ProductionRatesPiecePerDay int            `json:"production_rates_per_day,omitempty"`
	ThumbnailID                *uuid.UUID     `json:"thumbnail_id,omitempty"`
	Thumbnail                  *Images        `json:"thumbnail,omitempty" gorm:"foreignKey:ThumbnailID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	YearFounded                int            `json:"year_founded" gorm:"not null"`
	EmployeesCount             int            `json:"employees_count" gorm:"not null"`
	BusinessType               string         `json:"business_type" gorm:"not null"`
}

type CreateIndustriesBlogsReq struct {
	Title                      string     `json:"title" binding:"required"`
	Content                    string     `json:"content" binding:"required"`
	Location                   string     `json:"location"`
	Rating                     float64    `json:"rating"`
	Revenue                    float64    `json:"revenue"`
	ProducedProducts           []string   `json:"produced_products" binding:"required"`
	ProductionRatesPiecePerDay int        `json:"production_rates_per_day"`
	ThumbnailID                *uuid.UUID `json:"thumbnail_id"`
	YearFounded                int        `json:"year_founded" binding:"required"`
	EmployeesCount             int        `json:"employees_count" binding:"required"`
	BusinessType               string     `json:"business_type" binding:"required"`
}

type UpdateIndustriesBlogsReq struct {
	Title                      string     `json:"title"`
	Content                    string     `json:"content"`
	Location                   *string    `json:"location"`
	Rating                     *float64   `json:"rating"`
	Revenue                    *float64   `json:"revenue"`
	ProducedProducts           []string   `json:"produced_products" binding:"required"`
	ProductionRatesPiecePerDay *int       `json:"production_rates_per_day"`
	ThumbnailID                *uuid.UUID `json:"thumbnail_id"`
	YearFounded                *int       `json:"year_founded"`
	EmployeesCount             *int       `json:"employees_count"`
	BusinessType               *string    `json:"business_type"`
}
