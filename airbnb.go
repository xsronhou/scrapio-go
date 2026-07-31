package scrapio

import (
	"context"
	"fmt"
	"net/url"
)

type AirbnbSearchResponse struct {
	Provider     string         `json:"provider"`
	PageType     string         `json:"page_type"`
	SnapshotAt   string         `json:"snapshot_at"`
	QueryContext map[string]any `json:"query_context"`
	MetaData     map[string]any `json:"meta_data"`
	Results      []any          `json:"results"`
}

type AirbnbListingResponse struct {
	Provider   string         `json:"provider"`
	PageType   string         `json:"page_type"`
	SnapshotAt string         `json:"snapshot_at"`
	Listing    map[string]any `json:"listing"`
}

type AirbnbSearchRequest struct {
	Location string
	CheckIn  string
	CheckOut string
	Adults   int
	Children int
	Infants  int
	Pets     int
}

type AirbnbListingRequest struct {
	ListingID string
	CheckIn   string
	CheckOut  string
	Adults    int
}

type AirbnbResource struct{ h *httpClient }

func (r *AirbnbResource) Search(ctx context.Context, req AirbnbSearchRequest) (*AirbnbSearchResponse, error) {
	q := url.Values{"location": {req.Location}}
	if req.CheckIn != "" {
		q.Set("check_in", req.CheckIn)
	}
	if req.CheckOut != "" {
		q.Set("check_out", req.CheckOut)
	}
	if req.Adults > 0 {
		q.Set("adults", fmt.Sprintf("%d", req.Adults))
	}
	if req.Children > 0 {
		q.Set("children", fmt.Sprintf("%d", req.Children))
	}
	if req.Infants > 0 {
		q.Set("infants", fmt.Sprintf("%d", req.Infants))
	}
	if req.Pets > 0 {
		q.Set("pets", fmt.Sprintf("%d", req.Pets))
	}
	var out AirbnbSearchResponse
	return &out, r.h.get(ctx, "/v1/airbnb/search", q, &out)
}

func (r *AirbnbResource) GetListing(ctx context.Context, req AirbnbListingRequest) (*AirbnbListingResponse, error) {
	q := url.Values{"listing_id": {req.ListingID}}
	if req.CheckIn != "" {
		q.Set("check_in", req.CheckIn)
	}
	if req.CheckOut != "" {
		q.Set("check_out", req.CheckOut)
	}
	if req.Adults > 0 {
		q.Set("adults", fmt.Sprintf("%d", req.Adults))
	}
	var out AirbnbListingResponse
	return &out, r.h.get(ctx, "/v1/airbnb/listing", q, &out)
}
