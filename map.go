package scrapio

import "context"

// MapResource wraps POST /v1/map — site-map/URL-discovery workflow. Untyped
// request/response, same reasoning as SearchResource.
type MapResource struct{ h *httpClient }

func (r *MapResource) Execute(ctx context.Context, req map[string]any) (map[string]any, error) {
	var out map[string]any
	return out, r.h.post(ctx, "/v1/map", req, &out)
}
