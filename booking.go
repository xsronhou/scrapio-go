package scrapio

import (
	"context"
	"fmt"
	"net/url"
)

type HotelProperty struct {
	Position    int            `json:"position"`
	PropertyID  string         `json:"property_id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Type        string         `json:"type,omitempty"`
	Stars       *int           `json:"stars,omitempty"`
	ReviewScore *float64       `json:"review_score,omitempty"`
	ReviewCount *int           `json:"review_count,omitempty"`
	Price       map[string]any `json:"price"`
	Address     string         `json:"address,omitempty"`
	URL         string         `json:"url"`
}

type HotelSearchResponse struct {
	Provider       string          `json:"provider"`
	PageType       string          `json:"page_type"`
	SnapshotAt     string          `json:"snapshot_at"`
	MetaData       map[string]any  `json:"meta_data"`
	Stay           map[string]any  `json:"stay"`
	Occupancy      map[string]any  `json:"occupancy"`
	Currency       string          `json:"currency,omitempty"`
	PosCountry     string          `json:"pos_country,omitempty"`
	Properties     []HotelProperty `json:"properties"`
	AppliedFilters map[string]any  `json:"applied_filters"`
}

type HotelPropertyResponse struct {
	Provider    string         `json:"provider"`
	PageType    string         `json:"page_type"`
	SnapshotAt  string         `json:"snapshot_at"`
	PropertyID  string         `json:"property_id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Stars       *int           `json:"stars,omitempty"`
	Description string         `json:"description,omitempty"`
	Address     string         `json:"address,omitempty"`
	ReviewScore *float64       `json:"review_score,omitempty"`
	ReviewCount *int           `json:"review_count,omitempty"`
	Amenities   []string       `json:"amenities,omitempty"`
	Images      []string       `json:"images,omitempty"`
	Stay        map[string]any `json:"stay"`
	Occupancy   map[string]any `json:"occupancy"`
	Currency    string         `json:"currency,omitempty"`
	Available   bool           `json:"available"`
	Rooms       []any          `json:"rooms,omitempty"`
	URL         string         `json:"url"`
}

type HotelReviewsResponse struct {
	Provider       string             `json:"provider"`
	PageType       string             `json:"page_type"`
	SnapshotAt     string             `json:"snapshot_at"`
	PropertyID     string             `json:"property_id,omitempty"`
	Name           string             `json:"name,omitempty"`
	ReviewScore    *float64           `json:"review_score,omitempty"`
	ReviewCount    *int               `json:"review_count,omitempty"`
	MetaData       map[string]any     `json:"meta_data"`
	ScoreBreakdown map[string]float64 `json:"score_breakdown"`
	Reviews        []any              `json:"reviews"`
}

// HotelStayOpts carries the optional stay/occupancy/pricing parameters shared
// by Booking.com's and Agoda's search and property endpoints.
type HotelStayOpts struct {
	Adults       int
	Children     int
	ChildrenAges string
	Rooms        int
	Currency     string
	PosCountry   string
}

type HotelSearchOpts struct {
	HotelStayOpts
	SortBy         string
	MinPrice       float64
	MaxPrice       float64
	Stars          string
	MinReviewScore float64
	Page           int
}

func applyHotelStayOpts(q url.Values, o HotelStayOpts) {
	if o.Adults > 0 {
		q.Set("adults", fmt.Sprintf("%d", o.Adults))
	}
	if o.Children > 0 {
		q.Set("children", fmt.Sprintf("%d", o.Children))
	}
	if o.ChildrenAges != "" {
		q.Set("children_ages", o.ChildrenAges)
	}
	if o.Rooms > 0 {
		q.Set("rooms", fmt.Sprintf("%d", o.Rooms))
	}
	if o.Currency != "" {
		q.Set("currency", o.Currency)
	}
	if o.PosCountry != "" {
		q.Set("pos_country", o.PosCountry)
	}
}

func applyHotelSearchOpts(q url.Values, o HotelSearchOpts) {
	applyHotelStayOpts(q, o.HotelStayOpts)
	if o.SortBy != "" {
		q.Set("sort_by", o.SortBy)
	}
	if o.MinPrice > 0 {
		q.Set("min_price", fmt.Sprintf("%g", o.MinPrice))
	}
	if o.MaxPrice > 0 {
		q.Set("max_price", fmt.Sprintf("%g", o.MaxPrice))
	}
	if o.Stars != "" {
		q.Set("stars", o.Stars)
	}
	if o.MinReviewScore > 0 {
		q.Set("min_review_score", fmt.Sprintf("%g", o.MinReviewScore))
	}
	if o.Page > 0 {
		q.Set("page", fmt.Sprintf("%d", o.Page))
	}
}

// BookingPropertyRequest identifies a property by PropertyID or URL (exactly
// one should be set) plus the stay window and optional occupancy/pricing.
type BookingPropertyRequest struct {
	PropertyID string
	URL        string
	CheckIn    string
	CheckOut   string
	HotelStayOpts
}

// BookingReviewsRequest identifies a property by PropertyID or URL (exactly
// one should be set) plus optional paging/language/detail params.
type BookingReviewsRequest struct {
	PropertyID           string
	URL                  string
	Page                 int
	Language             string
	IncludeReviewDetails bool
}

type BookingResource struct{ h *httpClient }

func (r *BookingResource) Search(ctx context.Context, location, checkIn, checkOut string, opts ...HotelSearchOpts) (*HotelSearchResponse, error) {
	q := url.Values{"location": {location}, "check_in": {checkIn}, "check_out": {checkOut}}
	if len(opts) > 0 {
		applyHotelSearchOpts(q, opts[0])
	}
	var out HotelSearchResponse
	return &out, r.h.get(ctx, "/v1/booking/search", q, &out)
}

func (r *BookingResource) GetProperty(ctx context.Context, req BookingPropertyRequest) (*HotelPropertyResponse, error) {
	q := url.Values{"check_in": {req.CheckIn}, "check_out": {req.CheckOut}}
	if req.URL != "" {
		q.Set("url", req.URL)
	} else {
		q.Set("property_id", req.PropertyID)
	}
	applyHotelStayOpts(q, req.HotelStayOpts)
	var out HotelPropertyResponse
	return &out, r.h.get(ctx, "/v1/booking/property", q, &out)
}

func (r *BookingResource) GetReviews(ctx context.Context, req BookingReviewsRequest) (*HotelReviewsResponse, error) {
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
	if req.IncludeReviewDetails {
		q.Set("include_review_details", "true")
	}
	var out HotelReviewsResponse
	return &out, r.h.get(ctx, "/v1/booking/reviews", q, &out)
}
