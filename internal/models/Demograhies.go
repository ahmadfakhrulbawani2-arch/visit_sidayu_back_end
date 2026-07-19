package models

type Demographies struct {
	BaseModel
	VillageName            string `json:"village_name" gorm:"unique;not null"`
	DemographyDataYear     int    `json:"demography_data_year" gorm:"not null"`
	MalePopulation         int    `json:"male_population,omitempty"`
	FemalePopulation       int    `json:"female_population,omitempty"`
	TotalPopulation        int    `json:"total_population,omitempty"`
	PopulationDensityUnit  string `json:"population_density_unit" gorm:"not null"`
	FamiliesNumber         int    `json:"families_number,omitempty"`
	NumberOfBirth          int    `json:"number_of_birth,omitempty"`
	NumberOfDeath          int    `json:"number_of_death,omitempty"`
	WorkingPopulation      int    `json:"working_population,omitempty"`
	UnemployedPopulation   int    `json:"unemployed_population,omitempty"`
	HousekeepingPopulation int    `json:"housekeeping_population,omitempty"`
	StudentPopulation      int    `json:"student_population,omitempty"`
	SourceName             string `json:"source_name" gorm:"not null"`
	ExternalLinkSource     string `json:"external_link_source" gorm:"not null"`
}

type CreateDemographies struct {
	VillageName            string `json:"village_name" binding:"required"`
	DemographyDataYear     int    `json:"demography_data_year" binding:"required"`
	MalePopulation         int    `json:"male_population"`
	FemalePopulation       int    `json:"female_population"`
	TotalPopulation        int    `json:"total_population"`
	PopulationDensityUnit  string `json:"population_density_unit" binding:"required"`
	FamiliesNumber         int    `json:"families_number"`
	NumberOfBirth          int    `json:"number_of_birth"`
	NumberOfDeath          int    `json:"number_of_death"`
	WorkingPopulation      int    `json:"working_population"`
	UnemployedPopulation   int    `json:"unemployed_population"`
	HousekeepingPopulation int    `json:"housekeeping_population"`
	StudentPopulation      int    `json:"student_population"`
	SourceName             string `json:"source_name" binding:"required"`
	ExternalLinkSource     string `json:"external_link_source" binding:"required"`
}

type GetDistrictDemographies struct {
	DemographyDataYear    int      `json:"demography_data_year"`
	MalePopulation        int      `json:"male_population,omitempty"`
	FemalePopulation      int      `json:"female_population,omitempty"`
	TotalPopulation       int      `json:"total_population,omitempty"`
	PopulationDensity     float64  `json:"population_density"`
	PopulationDensityUnit string   `json:"population_density_unit"`
	Sources               []string `json:"sources"`
}
