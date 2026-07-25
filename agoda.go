package scrapio

import (
	"context"
	"fmt"
	"net/url"
)

// AgodaPropertyRequest identifies a property by URL, or by PropertyID
// together with CityID (from a prior Search's resolved city), plus the stay
// window and optional occupancy/pricing.
type AgodaPropertyRequest struct {
	PropertyID string
	CityID     int
	URL        string
	CheckIn    string
	CheckOut   string
	HotelStayOpts
}

// AgodaReviewsRequest identifies a property by PropertyID or URL (exactly
// one should be set) plus optional paging/language params.
type AgodaReviewsRequest struct {
	PropertyID string
	URL        string
	Page       int
	Language   string
}

type AgodaResource struct{ h *httpClient }

func (r *AgodaResource) Search(ctx context.Context, location, checkIn, checkOut string, opts ...HotelSearchOpts) (*HotelSearchResponse, error) {
	q := url.Values{"location": {location}, "check_in": {checkIn}, "check_out": {checkOut}}
	if len(opts) > 0 {
		applyHotelSearchOpts(q, opts[0])
	}
	var out HotelSearchResponse
	return &out, r.h.get(ctx, "/v1/agoda/search", q, &out)
}

func (r *AgodaResource) GetProperty(ctx context.Context, req AgodaPropertyRequest) (*HotelPropertyResponse, error) {
	q := url.Values{"check_in": {req.CheckIn}, "check_out": {req.CheckOut}}
	if req.URL != "" {
		q.Set("url", req.URL)
	} else {
		q.Set("property_id", req.PropertyID)
		q.Set("city_id", fmt.Sprintf("%d", req.CityID))
	}
	applyHotelStayOpts(q, req.HotelStayOpts)
	var out HotelPropertyResponse
	return &out, r.h.get(ctx, "/v1/agoda/property", q, &out)
}

func (r *AgodaResource) GetReviews(ctx context.Context, req AgodaReviewsRequest) (*HotelReviewsResponse, error) {
	q := url.Values{}
	if req.URL != "" {
		q.Set("url", req.URL)
	} else {
		q.Set("property_id", req.PropertyID)
	}
	if req.Page > 0 {
		q.Set("page", fmt.Sprintf("%d", req.Page))
	}
	if req.Language != "" {
		q.Set("language", req.Language)
	}
	var out HotelReviewsResponse
	return &out, r.h.get(ctx, "/v1/agoda/reviews", q, &out)
}
