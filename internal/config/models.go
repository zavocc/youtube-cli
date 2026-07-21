package config

import (
	"fmt"

	"google.golang.org/genai"
)

type configTemplate struct {
	ModelID        string
	ThinkingConfig *genai.ThinkingConfig
}

func ValidateModels(model string) (configTemplate, error) {
	switch model {
	case "gemini-3-flash-preview":
		return configTemplate{
			ModelID: "gemini-3-flash-preview",
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingLevel:   genai.ThinkingLevelMinimal,
				IncludeThoughts: false,
			},
		}, nil
	case "gemini-3.5-flash-lite":
		return configTemplate{
			ModelID: "gemini-3.5-flash-lite",
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingLevel:   genai.ThinkingLevelLow,
				IncludeThoughts: false,
			},
		}, nil
	case "gemini-3.6-flash":
		return configTemplate{
			ModelID: "gemini-3.6-flash",
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingLevel:   genai.ThinkingLevelMinimal,
				IncludeThoughts: false,
			},
		}, nil
	default:
		return configTemplate{}, fmt.Errorf("invalid model %q", model)
	}
}
