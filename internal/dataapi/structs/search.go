package structs

type CompactSearchResult struct {
	Type         string `json:"type"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	VideoID      string `json:"videoId"`
	ChannelID    string `json:"channelId"`
	ChannelTitle string `json:"channelTitle"`
	PublishedAt  string `json:"publishedAt"`
}

type CompactPlaylistsResult struct {
	Type         string `json:"type"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	PlaylistID   string `json:"playlistId"`
	ChannelID    string `json:"channelId"`
	ChannelTitle string `json:"channelTitle"`
	PublishedAt  string `json:"publishedAt"`
}

type CompactChannelResult struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ChannelID   string `json:"channelId"`
	PublishedAt string `json:"publishedAt"`
}
