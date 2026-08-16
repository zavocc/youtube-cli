package gemini

import (
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"github.com/OpenRouterTeam/go-sdk/optionalnullable"
	"google.golang.org/genai"
)

func genResponseSchema() *genai.Schema {
	return &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"answer": {
				Type:        genai.TypeString,
				Description: "Direct answer to the user's prompt about the video.",
			},
			"evidence_timestamps": {
				Type:        genai.TypeArray,
				Description: "Supporting timestamp and passages from the video supporting the answer.",
				Items: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"timestamp": {
							Type:        genai.TypeString,
							Description: "Video timestamp in MM:SS or HH:MM:SS format.",
						},
						"passage": {
							Type:        genai.TypeString,
							Description: "Short observation, quote, or paraphrase from that timestamp.",
						},
					},
					PropertyOrdering: []string{"timestamp", "passage"},
					Required:         []string{"timestamp", "passage"},
				},
			},
		},
		PropertyOrdering: []string{
			"answer",
			"evidence_timestamps",
		},
		Required: []string{
			"answer",
		},
	}
}

func genResponseSchemaOpenRouter() components.ResponseFormat {
	description := "Answer the prompt grounded from the video."
	strict := true

	return components.CreateResponseFormatJSONSchema(
		components.ChatFormatJSONSchemaConfig{
			JSONSchema: components.ChatJSONSchemaConfig{
				Name:        "answer_video",
				Description: &description,
				Strict:      optionalnullable.From(&strict),
				Schema: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"answer": map[string]any{
							"type":        "string",
							"description": "Direct answer to the user's prompt about the video.",
						},
						"evidence_timestamps": map[string]any{
							"type":        "array",
							"description": "Supporting timestamp and passages from the video supporting the answer.",
							"items": map[string]any{
								"type": "object",
								"properties": map[string]any{
									"timestamp": map[string]any{
										"type":        "string",
										"description": "Video timestamp in MM:SS or HH:MM:SS format.",
									},
									"passage": map[string]any{
										"type":        "string",
										"description": "Short observation, quote, or paraphrase from that timestamp.",
									},
								},
								"required": []string{
									"timestamp",
									"passage",
								},
							},
						},
					},
					"required": []string{
						"answer",
					},
				},
			},
		},
	)
}
