package structs

import (
	"encoding/json"
	"io"
	"strings"

	"google.golang.org/api/youtube/v3"
)

type CompactResponse struct {
	NextPageToken string `json:"nextPageToken"`
	Items         []any  `json:"items"`
}

// split whitespace delimited strings with string.Fields and join them with a space to consolidate and compact

func GenerateCompactJsonSearch(writer io.Writer, searchResponse *youtube.SearchListResponse) error {
	items := make([]any, 0, len(searchResponse.Items))

	for _, item := range searchResponse.Items {
		if item.Id == nil {
			continue
		}

		title := ""
		description := ""
		channelID := ""
		channelTitle := ""
		publishedAt := ""
		if item.Snippet != nil {
			title = strings.Join(strings.Fields(item.Snippet.Title), " ")
			description = strings.Join(strings.Fields(item.Snippet.Description), " ")
			channelID = item.Snippet.ChannelId
			channelTitle = strings.Join(strings.Fields(item.Snippet.ChannelTitle), " ")
			publishedAt = item.Snippet.PublishedAt
		}

		switch {
		case item.Id.VideoId != "":
			items = append(items, CompactSearchResult{
				Type:         "video",
				Title:        title,
				Description:  description,
				VideoID:      item.Id.VideoId,
				ChannelID:    channelID,
				ChannelTitle: channelTitle,
				PublishedAt:  publishedAt,
			})
		case item.Id.PlaylistId != "":
			items = append(items, CompactPlaylistsResult{
				Type:         "playlist",
				Title:        title,
				Description:  description,
				PlaylistID:   item.Id.PlaylistId,
				ChannelID:    channelID,
				ChannelTitle: channelTitle,
				PublishedAt:  publishedAt,
			})
		case item.Id.ChannelId != "":
			items = append(items, CompactChannelResult{
				Type:        "channel",
				Title:       title,
				Description: description,
				ChannelID:   item.Id.ChannelId,
				PublishedAt: publishedAt,
			})
		}
	}

	encoder := json.NewEncoder(writer)
	return encoder.Encode(CompactResponse{
		NextPageToken: searchResponse.NextPageToken,
		Items:         items,
	})
}

func GenerateCompactJsonPlaylists(writer io.Writer, playlistResponse *youtube.PlaylistItemListResponse) error {
	items := make([]any, 0, len(playlistResponse.Items))

	for _, item := range playlistResponse.Items {
		result := CompactPlaylistVideos{}

		if item.Snippet != nil {
			result.Title = strings.Join(strings.Fields(item.Snippet.Title), " ")
			result.Description = strings.Join(strings.Fields(item.Snippet.Description), " ")
			result.PlaylistID = item.Snippet.PlaylistId
			result.ChannelID = item.Snippet.ChannelId
			result.ChannelTitle = strings.Join(strings.Fields(item.Snippet.ChannelTitle), " ")
			result.PublishedAt = item.Snippet.PublishedAt
		}

		if item.ContentDetails != nil {
			result.VideoID = item.ContentDetails.VideoId
		}

		items = append(items, result)
	}

	encoder := json.NewEncoder(writer)
	return encoder.Encode(CompactResponse{
		NextPageToken: playlistResponse.NextPageToken,
		Items:         items,
	})
}
