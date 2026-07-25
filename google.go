package scrapio

import (
	"context"
	"fmt"
	"net/url"
)

type GoogleSearchParams struct {
	Search     string  `json:"search"`
	SearchType string  `json:"search_type,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Language   string  `json:"language,omitempty"`
	Device     string  `json:"device,omitempty"`
	Page       *int    `json:"page,omitempty"`
	DateRange  string  `json:"date_range,omitempty"`
	Latitude   *float64 `json:"latitude,omitempty"`
	Longitude  *float64 `json:"longitude,omitempty"`
	Radius     *float64 `json:"radius,omitempty"`
	SortBy     string  `json:"sort_by,omitempty"`
}

type GoogleSearchResponse struct {
	RequestID  string `json:"request_id"`
	Results    []any  `json:"results"`
	Pagination any    `json:"pagination,omitempty"`
}

type GoogleResource struct{ h *httpClient }

func (r *GoogleResource) Search(ctx context.Context, params *GoogleSearchParams) (*GoogleSearchResponse, error) {
	q := url.Values{"search": {params.Search}}
	if params.SearchType != "" {
		q.Set("search_type", params.SearchType)
	}
	if params.CountryCode != "" {
		q.Set("country_code", params.CountryCode)
	}
	if params.Language != "" {
		q.Set("language", params.Language)
	}
	if params.Device != "" {
		q.Set("device", params.Device)
	}
	if params.Page != nil {
		q.Set("page", fmt.Sprintf("%d", *params.Page))
	}
	if params.DateRange != "" {
		q.Set("date_range", params.DateRange)
	}
	var out GoogleSearchResponse
	return &out, r.h.get(ctx, "/v1/google/search", q, &out)
}
