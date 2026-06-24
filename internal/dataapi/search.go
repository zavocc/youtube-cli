package dataapi

import "google.golang.org/api/youtube/v3"

func Search(service *youtube.Service, query string, maxResults int64, nextPageToken string) (*youtube.SearchListResponse, error) {
	// Cap the maxResults to 50, if exceeds
	if maxResults > 50 {
		maxResults = 50
	}

	call := service.Search.List([]string{"id", "snippet"}).
		Q(query).
		Type("video").
		MaxResults(maxResults).
		SafeSearch("none")

	if nextPageToken != "" {
		call = call.PageToken(nextPageToken)
	}

	return call.Do()
}
