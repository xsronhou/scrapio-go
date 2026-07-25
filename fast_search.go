package scrapio

import (
	"context"
	"fmt"
	"net/url"
)

type FastSearchResponse struct {
	MetaData      map[string]any `json:"meta_data,omitempty"`
	OrganicResults []any         `json:"organic_results,omitempty"`
	NewsResults    []any         `json:"news_results,omitempty"`
}

type FastSearchResource struct{ h *httpClient }

func (r *FastSearchResource) Search(ctx context.Context, search string, opts ...FastSearchOpts) (*FastSearchResponse, error) {
	q := url.Values{"search": {search}}
	if len(opts) > 0 {
		if opts[0].CountryCode != "" {
			q.Set("country_code", opts[0].CountryCode)
		}
		if opts[0].Language != "" {
			q.Set("language", opts[0].Language)
		}
		if opts[0].Page > 0 {
			q.Set("page", fmt.Sprintf("%d", opts[0].Page))
		}
	}
	var out FastSearchResponse
	return &out, r.h.get(ctx, "/v1/fast-search", q, &out)
}

type FastSearchOpts struct {
	CountryCode string
	Language    string
	Page        int
}
