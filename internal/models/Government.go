package models

import "github.com/google/uuid"

/*
	Client must create role first and then create the officials
*/

type Roles struct {
	BaseModel
	Name        string `json:"name" gorm:"not null"`
	Level       int    `json:"level" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}

type Officials struct {
	BaseModel
	Name           string    `json:"name" gorm:"not null"`
	RoleID         uuid.UUID `json:"role_id"`
	Role           Roles     `json:"role" gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ProfileImageID uuid.UUID `json:"profile_image_id"`
	ProfileImage   Images    `json:"profile_image" gorm:"foreignKey:ProfileImageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Description    string    `json:"description"`
}

type CreateOfficials struct {
	Name           string    `json:"name" binding:"required"`
	Description    string    `json:"description"`
	RoleID         uuid.UUID `json:"role_id"`
	ProfileImageID uuid.UUID `json:"profile_image_id"`
}
