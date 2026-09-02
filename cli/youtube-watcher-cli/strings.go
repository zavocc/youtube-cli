package main

import "github.com/zavocc/youtube-watcher-cli/internal/gemini"

const (
	helpVideoString           = "YouTube video URL or ID [REQUIRED]"
	helpModelString           = "Model to use for inference, defaults to " + gemini.DefaultModel
	helpMediaResolutionString = "Media resolution for the video. Possible values are only low, high. If not set, it will default for low resolution (this only works if not using --openrouter flag)"
	helpServiceTierString     = "Service tier. Possible values are 'flex' or 'priority'. Leave this for standard rates processing."
	helpOpenRouterString      = "Use OpenRouter endpoint instead of Google, note you need to set an OpenRouter API key using same environment variable instead of Generative Client API."
	helpUseAgenticProcessing  = "Use agentic processing of videos instead of ingesting each video frames"
	helpPromptString          = "Prompt to ask questions about the video [REQUIRED]"
	helpShowHelpString        = "Show help"
	helpVersionString         = "Print version"
)
