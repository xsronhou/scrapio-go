package scrapio

import (
	"context"
	"fmt"
	"net/url"
)

type WalmartProductResponse struct {
	Provider     string   `json:"provider"`
	ProductID    string   `json:"product_id"`
	Title        string   `json:"title"`
	Brand        string   `json:"brand,omitempty"`
	Price        *float64 `json:"price,omitempty"`
	Availability string   `json:"availability,omitempty"`
	URL          string   `json:"url"`
}

type WalmartSearchResponse struct {
	RequestID string                   `json:"request_id"`
	Results   []WalmartProductResponse `json:"results"`
}

type WalmartResource struct{ h *httpClient }

// GetProduct calls GET /v1/walmart/product.
//
// Deprecated: /v1/walmart/product is temporarily disabled server-side as of
// 2026-07-11 (frequent PerimeterX-related blocks on individual Walmart
// product pages) and will return a 404 until it's re-enabled. Search is
// unaffected. See specs/product/walmart-api.md in the main repo.
func (r *WalmartResource) GetProduct(ctx context.Context, productID string, zipCode ...string) (*WalmartProductResponse, error) {
	q := url.Values{"product_id": {productID}}
	if len(zipCode) > 0 && zipCode[0] != "" {
		q.Set("zip_code", zipCode[0])
	}
	var out WalmartProductResponse
	return &out, r.h.get(ctx, "/v1/walmart/product", q, &out)
}

func (r *WalmartResource) Search(ctx context.Context, search string, opts ...WalmartSearchOpts) (*WalmartSearchResponse, error) {
	q := url.Values{"search": {search}}
	if len(opts) > 0 {
		if opts[0].ZipCode != "" {
			q.Set("zip_code", opts[0].ZipCode)
		}
		if opts[0].Page > 0 {
			q.Set("page", fmt.Sprintf("%d", opts[0].Page))
		}
	}
	var out WalmartSearchResponse
	return &out, r.h.get(ctx, "/v1/walmart/search", q, &out)
}

type WalmartSearchOpts struct {
	ZipCode string
	Page    int
}
