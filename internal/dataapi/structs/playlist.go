package structs

type CompactPlaylistVideos struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	PlaylistID   string `json:"playlistId"`
	VideoID      string `json:"videoId"`
	ChannelID    string `json:"channelId"`
	ChannelTitle string `json:"channelTitle"`
	PublishedAt  string `json:"publishedAt"`
}
