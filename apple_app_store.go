package scrapio

import (
	"context"
	"net/url"
)

type AppleAppStoreRating struct {
	Value *float64 `json:"value"`
	Count *int     `json:"count"`
}

type AppleAppStoreAuthor struct {
	Name *string `json:"name"`
	URL  *string `json:"url"`
}

type AppleAppStoreApp struct {
	URL                string              `json:"url"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	ImageURL           *string             `json:"image_url"`
	Price              *float64            `json:"price"`
	Currency           *string             `json:"currency"`
	Category           *string             `json:"category"`
	SubCategory        *string             `json:"sub_category"`
	Rating             AppleAppStoreRating `json:"rating"`
	Author             AppleAppStoreAuthor `json:"author"`
	OperatingSystem    *string             `json:"operating_system"`
	AvailableOnDevice  *string             `json:"available_on_device"`
}

type AppleAppStoreResponse struct {
	Provider  string           `json:"provider"`
	RequestID string           `json:"request_id"`
	App       AppleAppStoreApp `json:"app"`
}

type AppleAppStoreResource struct{ h *httpClient }

func (r *AppleAppStoreResource) GetApp(ctx context.Context, appURL string) (*AppleAppStoreResponse, error) {
	q := url.Values{"url": {appURL}}
	var out AppleAppStoreResponse
	return &out, r.h.get(ctx, "/v1/apple-app-store/app", q, &out)
}
