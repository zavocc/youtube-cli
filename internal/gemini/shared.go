package gemini

import (
	"fmt"
	"strings"
)

func checkUrl(url string) (string, error) {
	// Check if it's either 11-character YouTube video ID or a full URL
	var actualUrl string
	if len(url) == 11 {
		actualUrl = "https://www.youtube.com/watch?v=" + url
	} else if len(url) > 11 && (url[:7] == "http://" || url[:8] == "https://") {
		actualUrl = strings.Split(url, "&")[0] // Remove any additional query parameters after the video ID
	} else {
		return "", fmt.Errorf("invalid YouTube video ID or URL specified")
	}
	return actualUrl, nil
}
