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

type YouTubeSearchOpts struct {
	StartPage int
	EndPage   int
	Country   string
	Language  string
}
