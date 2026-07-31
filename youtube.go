package scrapio

import (
	"context"
	"fmt"
	"net/url"
)

type YouTubeVideoResponse struct {
	RequestID string         `json:"request_id"`
	Video     map[string]any `json:"video"`
}

type YouTubeSearchResponse struct {
	RequestID string `json:"request_id"`
	Results   []any  `json:"results"`
}

type YouTubeSubtitleResponse struct {
	RequestID string `json:"request_id"`
	VideoID   string `json:"video_id"`
	Subtitles []any  `json:"subtitles"`
}

type YouTubeChannelLink struct {
	URL   string  `json:"url"`
	Title *string `json:"title"`
}

type YouTubeChannelVideo struct {
	Position   int     `json:"position"`
	URL        string  `json:"url"`
	Title      string  `json:"title"`
	ViewCount  *string `json:"view_count,omitempty"`
	UploadedAt *string `json:"uploaded_at,omitempty"`
}

type YouTubeChannel struct {
	ID               *string               `json:"id"`
	Name             *string               `json:"name"`
	Handle           string                `json:"handle"`
	Description      *string               `json:"description"`
	AvatarURL        *string               `json:"avatar_url"`
	Country          *string               `json:"country"`
	JoinedDate       *string               `json:"joined_date"`
	SubscriberCount  *string               `json:"subscriber_count"`
	ViewCount        *string               `json:"view_count"`
	VideoCount       *int                  `json:"video_count"`
	Links            []YouTubeChannelLink  `json:"links"`
	Videos           []YouTubeChannelVideo `json:"videos"`
}

type YouTubeChannelResponse struct {
	RequestID string         `json:"request_id"`
	Channel   YouTubeChannel `json:"channel"`
}

type YouTubeResource struct{ h *httpClient }

func (r *YouTubeResource) GetVideo(ctx context.Context, videoID string) (*YouTubeVideoResponse, error) {
	var out YouTubeVideoResponse
	return &out, r.h.get(ctx, "/v1/youtube/videos/"+videoID, nil, &out)
}

func (r *YouTubeResource) Search(ctx context.Context, search string, opts ...YouTubeSearchOpts) (*YouTubeSearchResponse, error) {
	q := url.Values{"search": {search}}
	if len(opts) > 0 {
		if opts[0].StartPage > 0 {
			q.Set("start_page", fmt.Sprintf("%d", opts[0].StartPage))
		}
		if opts[0].EndPage > 0 {
			q.Set("end_page", fmt.Sprintf("%d", opts[0].EndPage))
		}
		if opts[0].Country != "" {
			q.Set("country", opts[0].Country)
		}
		if opts[0].Language != "" {
			q.Set("language", opts[0].Language)
		}
	}
	var out YouTubeSearchResponse
	return &out, r.h.get(ctx, "/v1/youtube/search", q, &out)
}

func (r *YouTubeResource) GetSubtitles(ctx context.Context, videoID string, language ...string) (*YouTubeSubtitleResponse, error) {
	q := url.Values{"video_id": {videoID}}
	if len(language) > 0 && language[0] != "" {
		q.Set("language", language[0])
	}
	var out YouTubeSubtitleResponse
	return &out, r.h.get(ctx, "/v1/youtube/subtitles", q, &out)
}

func (r *YouTubeResource) GetChannel(ctx context.Context, handle string, limit ...int) (*YouTubeChannelResponse, error) {
	q := url.Values{"handle": {handle}}
	if len(limit) > 0 && limit[0] > 0 {
		q.Set("limit", fmt.Sprintf("%d", limit[0]))
	}
	var out YouTubeChannelResponse
	return &out, r.h.get(ctx, "/v1/youtube/channel", q, &out)
}

type YouTubeSearchOpts struct {
	StartPage int
	EndPage   int
	Country   string
	Language  string
}
