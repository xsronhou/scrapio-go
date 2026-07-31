package scrapio

import (
	"context"
	"net/url"
)

type GeminiPromptResponse struct {
	Provider     string            `json:"provider"`
	RequestID    string            `json:"request_id"`
	Prompt       string            `json:"prompt"`
	Model        *string           `json:"model"`
	ResponseText *string           `json:"response_text"`
	MarkdownText *string           `json:"markdown_text"`
	Citations    []ChatgptCitation `json:"citations"`
}

type GeminiPromptRequest struct {
	Prompt string
	Geo    string
}

type GeminiResource struct{ h *httpClient }

func (r *GeminiResource) Prompt(ctx context.Context, req GeminiPromptRequest) (*GeminiPromptResponse, error) {
	q := url.Values{"prompt": {req.Prompt}}
	if req.Geo != "" {
		q.Set("geo", req.Geo)
	}
	var out GeminiPromptResponse
	return &out, r.h.get(ctx, "/v1/gemini/prompt", q, &out)
}
