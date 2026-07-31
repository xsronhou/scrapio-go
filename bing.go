package scrapio

import (
	"context"
	"fmt"
	"net/url"
)

type BingResult struct {
	Position        int     `json:"position"`
	PositionOverall int     `json:"position_overall"`
	Title           string  `json:"title"`
	Description     *string `json:"description"`
	URL             string  `json:"url"`
	URLShown        *string `json:"url_shown"`
}

type BingSearchResponse struct {
	Provider       string         `json:"provider"`
	RequestID      string         `json:"request_id"`
	QueryContext   map[string]any `json:"query_context"`
	MetaData       map[string]any `json:"meta_data"`
	OrganicResults []BingResult   `json:"organic_results"`
	PaidResults    []BingResult   `json:"paid_results"`
}

type BingSearchRequest struct {
	Search    string
	Geo       string
	Locale    string
	Domain    string
	Page      int
	PageCount int
}

type BingResource struct{ h *httpClient }

func (r *BingResource) Search(ctx context.Context, req BingSearchRequest) (*BingSearchResponse, error) {
	q := url.Values{"search": {req.Search}}
	if req.Geo != "" {
		q.Set("geo", req.Geo)
	}
	if req.Locale != "" {
		q.Set("locale", req.Locale)
	}
	if req.Domain != "" {
		q.Set("domain", req.Domain)
	}
	if req.Page > 0 {
		q.Set("page", fmt.Sprintf("%d", req.Page))
	}
	if req.PageCount > 0 {
		q.Set("page_count", fmt.Sprintf("%d", req.PageCount))
	}
	var out BingSearchResponse
	return &out, r.h.get(ctx, "/v1/bing/search", q, &out)
}
