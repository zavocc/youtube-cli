package gemini

import (
	"context"
	"fmt"
	"os"
	"strings"

	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
	"github.com/OpenRouterTeam/go-sdk/optionalnullable"
	"github.com/zavocc/youtube-watcher-cli/internal/config"
)

func ORouterClient(ctx context.Context, prompt string, url string, model string, serviceTier string) (string, error) {
	// create new OpenRouter client
	if os.Getenv("OPENROUTER_API_KEY") == "" {
		return "", fmt.Errorf("OPENROUTER_API_KEY environment variable is not set")
	}

	orClient := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	// Check if it's either 11-character YouTube video ID or a full URL
	var actualUrl string
	if url, err := checkUrl(url); err != nil {
		return "", err
	} else {
		actualUrl = url
	}

	// Check if service tier is flex, priority, or standard
	var parsedServiceTier components.ChatRequestServiceTier
	switch strings.ToLower(strings.TrimSpace(serviceTier)) {
	case "flex":
		parsedServiceTier = components.ChatRequestServiceTierFlex
	case "priority":
		parsedServiceTier = components.ChatRequestServiceTierPriority
	case "standard":
		parsedServiceTier = components.ChatRequestServiceTierDefault
	default:
		return "", fmt.Errorf("invalid service tier specified. Must be 'flex', 'priority', or 'standard'")
	}

	// Validate model and get corresponding config
	modelSelectedConfig, err := config.ValidateModelsOpenRouter(model)
	if err != nil {
		return "", err
	}

	// Turns
	contents := []components.ChatMessages{
		components.CreateChatMessagesSystem(
			components.ChatSystemMessage{
				Content: components.CreateChatSystemMessageContentStr(systemPrompt),
				Role:    components.ChatSystemMessageRoleSystem,
			},
		),
		components.CreateChatMessagesUser(
			components.ChatUserMessage{
				Content: components.CreateChatUserMessageContentArrayOfChatContentItems(
					[]components.ChatContentItems{
						components.CreateChatContentItemsVideoURL(
							components.ChatContentVideo{
								VideoURL: components.ChatContentVideoInput{
									URL: actualUrl,
								},
							},
						),
					},
				),
				Role: components.ChatUserMessageRoleUser,
			},
		),
		components.CreateChatMessagesUser(
			components.ChatUserMessage{
				Content: components.CreateChatUserMessageContentStr(prompt),
				Role:    components.ChatUserMessageRoleUser,
			},
		),
	}

	// Create the chat request
	responseFormat := genResponseSchemaOpenRouter()
	temperatureNullable := 0.0
	allowProviderFallbacks := false
	onlyProviders := []components.ProviderPreferencesOnly{
		components.CreateProviderPreferencesOnlyStr("google-ai-studio"),
	}
	providerPreferences := components.ProviderPreferences{
		AllowFallbacks: optionalnullable.From(&allowProviderFallbacks),
		Only:           optionalnullable.From(&onlyProviders),
	}
	result, err := orClient.Chat.Send(ctx, components.ChatRequest{
		Messages:       contents,
		Model:          &modelSelectedConfig.ModelID,
		Provider:       optionalnullable.From(&providerPreferences),
		Reasoning:      modelSelectedConfig.ThinkingConfig,
		ServiceTier:    optionalnullable.From(&parsedServiceTier),
		ResponseFormat: &responseFormat,
		Temperature:    optionalnullable.From(&temperatureNullable),
	}, nil)

	if err != nil {
		return "", fmt.Errorf("perform openrouter inference: %w", err)
	}

	if result.ChatResult == nil {
		return "", fmt.Errorf("openrouter returned no chat result")
	}

	choices := result.ChatResult.GetChoices()
	if len(choices) == 0 {
		return "", fmt.Errorf("openrouter returned no choices")
	}

	content, isSet := choices[0].Message.Content.GetOrZero()
	if !isSet || content.Str == nil {
		return "", fmt.Errorf("openrouter returned no text content")
	}

	return *content.Str, nil
}
