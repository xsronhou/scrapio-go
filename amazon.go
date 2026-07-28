package scrapio

import (
	"context"
	"fmt"
	"net/url"
)

type AmazonProductResponse struct {
	Provider      string   `json:"provider"`
	ASIN          string   `json:"asin"`
	Title         string   `json:"title"`
	Brand         string   `json:"brand,omitempty"`
	Price         *float64 `json:"price,omitempty"`
	Currency      string   `json:"currency,omitempty"`
	Availability  string   `json:"availability,omitempty"`
	Rating        *float64 `json:"rating,omitempty"`
	ReviewCount   *int     `json:"review_count,omitempty"`
	Images        []string `json:"images,omitempty"`
	BulletPoints  []string `json:"bullet_points,omitempty"`
	URL           string   `json:"url"`
}

type AmazonSearchResponse struct {
	RequestID string                  `json:"request_id"`
	Results   []AmazonProductResponse `json:"results"`
}

type AmazonResource struct{ h *httpClient }

func (r *AmazonResource) GetProduct(ctx context.Context, asin string, country ...string) (*AmazonProductResponse, error) {
	q := url.Values{"asin": {asin}}
	if len(country) > 0 && country[0] != "" {
		q.Set("country", country[0])
	}
	var out AmazonProductResponse
	return &out, r.h.get(ctx, "/v1/amazon/product", q, &out)
}

func (r *AmazonResource) Search(ctx context.Context, search string, opts ...AmazonSearchOpts) (*AmazonSearchResponse, error) {
	q := url.Values{"search": {search}}
	if len(opts) > 0 {
		if opts[0].Country != "" {
			q.Set("country", opts[0].Country)
		}
		if opts[0].Page > 0 {
			q.Set("page", fmt.Sprintf("%d", opts[0].Page))
		}
	}
	var out AmazonSearchResponse
	return &out, r.h.get(ctx, "/v1/amazon/search", q, &out)
}

type AmazonSearchOpts struct {
	Country string
	Page    int
}
