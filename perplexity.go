package scrapio

import (
	"context"
	"net/url"
)

type PerplexitySource struct {
	URL    string  `json:"url"`
	Title  *string `json:"title"`
	Source *string `json:"source"`
}

type PerplexityPromptResponse struct {
	Provider       string             `json:"provider"`
	RequestID      string             `json:"request_id"`
	Prompt         string             `json:"prompt"`
	Model          *string            `json:"model"`
	AnswerMarkdown *string            `json:"answer_markdown"`
	RelatedQueries []string           `json:"related_queries"`
	Sources        []PerplexitySource `json:"sources"`
}

type PerplexityPromptRequest struct {
	Prompt string
	Geo    string
}

type PerplexityResource struct{ h *httpClient }

func (r *PerplexityResource) Prompt(ctx context.Context, req PerplexityPromptRequest) (*PerplexityPromptResponse, error) {
	q := url.Values{"prompt": {req.Prompt}}
	if req.Geo != "" {
		q.Set("geo", req.Geo)
	}
	var out PerplexityPromptResponse
	return &out, r.h.get(ctx, "/v1/perplexity/prompt", q, &out)
}
