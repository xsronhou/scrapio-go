package scrapio

import (
	"context"
	"net/url"
)

type ChatgptCitation struct {
	URL         string  `json:"url"`
	Text        *string `json:"text"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type ChatgptPromptResponse struct {
	Provider      string            `json:"provider"`
	RequestID     string            `json:"request_id"`
	Prompt        string            `json:"prompt"`
	Model         *string           `json:"model"`
	ResponseText  *string           `json:"response_text"`
	MarkdownText  *string           `json:"markdown_text"`
	SearchQueries []string          `json:"search_queries"`
	Citations     []ChatgptCitation `json:"citations"`
}

type ChatgptPromptRequest struct {
	Prompt string
	Search bool
	Geo    string
}

type ChatgptResource struct{ h *httpClient }

func (r *ChatgptResource) Prompt(ctx context.Context, req ChatgptPromptRequest) (*ChatgptPromptResponse, error) {
	q := url.Values{"prompt": {req.Prompt}}
	if req.Search {
		q.Set("search", "true")
	}
	if req.Geo != "" {
		q.Set("geo", req.Geo)
	}
	var out ChatgptPromptResponse
	return &out, r.h.get(ctx, "/v1/chatgpt/prompt", q, &out)
}
