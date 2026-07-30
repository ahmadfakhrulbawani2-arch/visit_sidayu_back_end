package helpers

import (
	"errors"
	"slices"
	"visit-sidayu-backend/internal/constants/input"
)

// population density, unit, error
func CalcPopulationDensity(population int, area float64, area_unit string) (float64, *string, error) {
	// meaning area data is not provided, thus this data can't be computed
	if area_unit == "" {
		return 0, nil, nil
	}

	if !slices.Contains(input.Area_Units, area_unit) {
		return 0, nil, errors.New("invalid area unit")
	}

	if area == 0 {
		return 0, nil, nil
	}

	unit := "jiwa/" + area_unit

	return (float64(population) / float64(area)), &unit, nil
}
