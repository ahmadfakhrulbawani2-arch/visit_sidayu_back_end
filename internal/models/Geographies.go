package models

import "github.com/google/uuid"

// assume we don't want to miss the unit measurement so even if the data is unknwon we still fill the unit
type Geographies struct {
	BaseModel
	VillageName  string     `json:"village_name" gorm:"unique;not null"` // the name should be unique as we make only 1 district
	Area         float64    `json:"area,omitempty"`
	AreaUnit     string     `json:"area_unit" gorm:"not null"`
	RainfallRate float64    `json:"rainfall_rate,omitempty"`
	RainfallUnit string     `json:"rainfall_unit" gorm:"not null"`
	RainyDay     int        `json:"rainy_day,omitempty"`
	ImageID      *uuid.UUID `json:"image_id,omitempty"`
	Image        *Images    `json:"image,omitempty" gorm:"foreignKey:ImageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Source       string     `json:"source" gorm:"not null"`
}

type CreateGeographies struct {
	VillageName  string     `json:"village_name" binding:"required"`
	Area         float64    `json:"area"`
	AreaUnit     string     `json:"area_unit" binding:"required"`
	RainfallRate float64    `json:"rainfall_rate"`
	RainfallUnit string     `json:"rainfall_unit" binding:"required"`
	RainyDay     int        `json:"rainy_day"`
	ImageID      *uuid.UUID `json:"image_id"`
	Source       string     `json:"source" binding:"required"`
}

type GetDistrictGeographies struct {
	VillageCount        int      `json:"village_count"`
	AreaUnit            string   `json:"area_unit"`
	TotalArea           float64  `json:"total_area,omitempty"`
	SmallestVillageName string   `json:"smallest_village_name,omitempty"`
	LargestVillageName  string   `json:"largest_village_name,omitempty"`
	SmallestVillageArea float64  `json:"smallest_village_area,omitempty"`
	LargestVillageArea  float64  `json:"largest_village_area,omitempty"`
	Sources             []string `json:"sources"`
}
