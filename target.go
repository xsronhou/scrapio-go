package scrapio

import (
	"context"
	"net/url"
)

type TargetProductResponse struct {
	Provider  string         `json:"provider"`
	RequestID string         `json:"request_id"`
	Product   map[string]any `json:"product"`
}

type TargetSearchResponse struct {
	Provider     string         `json:"provider"`
	RequestID    string         `json:"request_id"`
	QueryContext map[string]any `json:"query_context"`
	Results      []any          `json:"results"`
}

type TargetDeliveryOpts struct {
	DeliveryType string
	DeliveryZip  string
	StoreID      string
}

func applyTargetDeliveryOpts(q url.Values, opts TargetDeliveryOpts) {
	if opts.DeliveryType != "" {
		q.Set("delivery_type", opts.DeliveryType)
	}
	if opts.DeliveryZip != "" {
		q.Set("delivery_zip", opts.DeliveryZip)
	}
	if opts.StoreID != "" {
		q.Set("store_id", opts.StoreID)
	}
}

type TargetResource struct{ h *httpClient }

func (r *TargetResource) GetProduct(ctx context.Context, productID string, opts ...TargetDeliveryOpts) (*TargetProductResponse, error) {
	q := url.Values{"product_id": {productID}}
	if len(opts) > 0 {
		applyTargetDeliveryOpts(q, opts[0])
	}
	var out TargetProductResponse
	return &out, r.h.get(ctx, "/v1/target/product", q, &out)
}

func (r *TargetResource) Search(ctx context.Context, search string, opts ...TargetDeliveryOpts) (*TargetSearchResponse, error) {
	q := url.Values{"search": {search}}
	if len(opts) > 0 {
		applyTargetDeliveryOpts(q, opts[0])
	}
	var out TargetSearchResponse
	return &out, r.h.get(ctx, "/v1/target/search", q, &out)
}
