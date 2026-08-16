package config

import (
	"fmt"

	"github.com/OpenRouterTeam/go-sdk/models/components"
	"github.com/OpenRouterTeam/go-sdk/optionalnullable"
	"google.golang.org/genai"
)

type configTemplate struct {
	ModelID        string
	ThinkingConfig *genai.ThinkingConfig
}

type configTemplateOpenRouter struct {
	ModelID        string
	ThinkingConfig *components.ChatRequestReasoning
}

func chatRequestEffort(effort components.ChatRequestEffort) optionalnullable.OptionalNullable[components.ChatRequestEffort] {
	return optionalnullable.From(&effort)
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
	case "gemini-3.7-flash":
		return configTemplate{
			ModelID: "gemini-3.7-flash",
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingLevel:   genai.ThinkingLevelLow,
				IncludeThoughts: false,
			},
		}, nil
	default:
		return configTemplate{}, fmt.Errorf("invalid model %q", model)
	}
}

func ValidateModelsOpenRouter(model string) (configTemplateOpenRouter, error) {
	switch model {
	case "gemini-3-flash-preview":
		return configTemplateOpenRouter{
			ModelID: "gemini-3-flash-preview",
			ThinkingConfig: &components.ChatRequestReasoning{
				Effort: chatRequestEffort(components.ChatRequestEffortMinimal),
			},
		}, nil
	case "gemini-3.5-flash-lite":
		return configTemplateOpenRouter{
			ModelID: "gemini-3.5-flash-lite",
			ThinkingConfig: &components.ChatRequestReasoning{
				Effort: chatRequestEffort(components.ChatRequestEffortLow),
			},
		}, nil
	case "gemini-3.7-flash":
		return configTemplateOpenRouter{
			ModelID: "gemini-3.7-flash",
			ThinkingConfig: &components.ChatRequestReasoning{
				Effort: chatRequestEffort(components.ChatRequestEffortLow),
			},
		}, nil
	default:
		return configTemplateOpenRouter{}, fmt.Errorf("invalid model %q", model)
	}
}
