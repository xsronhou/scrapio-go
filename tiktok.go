package scrapio

import (
	"context"
	"net/url"
)

type TikTokPostAuthor struct {
	Username  *string `json:"username"`
	Nickname  *string `json:"nickname"`
	AvatarURL *string `json:"avatar_url"`
}

type TikTokPostStats struct {
	Likes    *int `json:"likes"`
	Comments *int `json:"comments"`
	Shares   *int `json:"shares"`
	Plays    *int `json:"plays"`
	Collects *int `json:"collects"`
}

type TikTokPostVideo struct {
	CoverURL *string `json:"cover_url"`
	PlayURL  *string `json:"play_url"`
}

type TikTokPostMusic struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
}

type TikTokPost struct {
	ID              string           `json:"id"`
	URL             string           `json:"url"`
	Description     *string          `json:"description"`
	CreatedAt       *string          `json:"created_at"`
	DurationSeconds *float64         `json:"duration_seconds"`
	Author          TikTokPostAuthor `json:"author"`
	Stats           TikTokPostStats  `json:"stats"`
	Video           TikTokPostVideo  `json:"video"`
	Music           TikTokPostMusic  `json:"music"`
}

type TikTokPostResponse struct {
	Provider  string     `json:"provider"`
	RequestID string     `json:"request_id"`
	Post      TikTokPost `json:"post"`
}

type TikTokResource struct{ h *httpClient }

func (r *TikTokResource) GetPost(ctx context.Context, postURL string) (*TikTokPostResponse, error) {
	q := url.Values{"url": {postURL}}
	var out TikTokPostResponse
	return &out, r.h.get(ctx, "/v1/tiktok/post", q, &out)
}
