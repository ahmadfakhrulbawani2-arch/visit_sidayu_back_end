package helpers

import (
	"errors"
	"slices"
	"visit-sidayu-backend/internal/constants/input"
)

func CalcPopulationDensity(population int, area float64, area_unit string) (float64, string, error) {
	if !slices.Contains(input.Area_Units, area_unit) {
		return 0, "-", errors.New("invalid area unit")
	}

	if area == 0 {
		return 0, "-", errors.New("area can not be zero")
	}

	return (float64(population) / float64(area)), "jiwa/" + area_unit, nil
}
