package scrapio

import (
	"context"
	"net/url"
)

type RedditPostResponse struct {
	Provider  string         `json:"provider"`
	RequestID string         `json:"request_id"`
	Post      map[string]any `json:"post"`
	Comments  []any          `json:"comments"`
}

type RedditSubredditResponse struct {
	Provider  string  `json:"provider"`
	RequestID string  `json:"request_id"`
	Subreddit *string `json:"subreddit"`
	Posts     []any   `json:"posts"`
}

type RedditUserResponse struct {
	Provider  string  `json:"provider"`
	RequestID string  `json:"request_id"`
	Username  *string `json:"username"`
	Items     []any   `json:"items"`
}

type RedditResource struct{ h *httpClient }

func (r *RedditResource) GetPost(ctx context.Context, postURL string) (*RedditPostResponse, error) {
	q := url.Values{"url": {postURL}}
	var out RedditPostResponse
	return &out, r.h.get(ctx, "/v1/reddit/post", q, &out)
}

func (r *RedditResource) GetSubreddit(ctx context.Context, subreddit, subredditURL string) (*RedditSubredditResponse, error) {
	q := url.Values{}
	if subreddit != "" {
		q.Set("subreddit", subreddit)
	}
	if subredditURL != "" {
		q.Set("url", subredditURL)
	}
	var out RedditSubredditResponse
	return &out, r.h.get(ctx, "/v1/reddit/subreddit", q, &out)
}

func (r *RedditResource) GetUser(ctx context.Context, username, userURL string) (*RedditUserResponse, error) {
	q := url.Values{}
	if username != "" {
		q.Set("username", username)
	}
	if userURL != "" {
		q.Set("url", userURL)
	}
	var out RedditUserResponse
	return &out, r.h.get(ctx, "/v1/reddit/user", q, &out)
}
