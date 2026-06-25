package dataapi

import (
	"fmt"

	"google.golang.org/api/youtube/v3"
)

func Channel(service *youtube.Service, query string, channelQueryType string, maxResults int64, nextPageToken string) (*youtube.ChannelListResponse, error) {
	// Cap the maxResults to 50, if exceeds
	if maxResults > 50 {
		maxResults = 50
	}

	call := service.Channels.List([]string{"id", "snippet", "contentDetails"}).
		MaxResults(maxResults)

	// check if channelQueryType is "username", "handle" or "id" and set the appropriate parameter, by default, it will be "handle"
	switch channelQueryType {
	case "id":
		call = call.Id(query)
	case "username":
		call = call.ForUsername(query)
	case "handle":
		call = call.ForHandle(query)
	default:
		return nil, fmt.Errorf("invalid channel query type %q: expected id, username or handle", channelQueryType)
	}

	if nextPageToken != "" {
		call = call.PageToken(nextPageToken)
	}

	return call.Do()
}
