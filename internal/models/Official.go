package models

import "github.com/google/uuid"

/*
	Client must create role first and then create the officials
*/

type Officials struct {
	BaseModel
	Name           string     `json:"name" gorm:"unique;not null"`
	Description    string     `json:"description,omitempty"`
	RoleID         *uuid.UUID `json:"role_id" gorm:"not null"`
	Role           *Roles     `json:"role" gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ProfileImageID *uuid.UUID `json:"profile_image_id,omitempty"`
	ProfileImage   *Images    `json:"profile_image,omitempty" gorm:"foreignKey:ProfileImageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreateOfficials struct {
	Name           string     `json:"name" binding:"required"`
	Description    string     `json:"description"`
	RoleID         *uuid.UUID `json:"role_id" binding:"required"`
	ProfileImageID *uuid.UUID `json:"profile_image_id"`
}

type UpdateOfficials struct {
	Name           string     `json:"name" binding:"required"`
	Description    *string    `json:"description"`
	RoleID         *uuid.UUID `json:"role_id" binding:"required"`
	ProfileImageID *uuid.UUID `json:"profile_image_id"`
}
