package config

import (
	"fmt"

	"google.golang.org/genai"
)

func ParseMediaResolution(resolution string) (genai.PartMediaResolutionLevel, error) {
	switch resolution {
	case "low":
		return genai.PartMediaResolutionLevelMediaResolutionLow, nil
	case "high":
		return genai.PartMediaResolutionLevelMediaResolutionHigh, nil
	default:
		return genai.PartMediaResolutionLevelMediaResolutionUnspecified, fmt.Errorf("invalid media resolution %q: expected low or high", resolution)
	}
}
