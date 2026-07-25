package scrapio

import "context"

// SearchResource wraps POST /v1/search — a multi-provider search workflow
// (currently resolves through Google, requires a Pro plan or higher).
// Request/response are intentionally untyped maps since the workflow
// orchestrator's payload is more flexible than a fixed schema.
type SearchResource struct{ h *httpClient }

func (r *SearchResource) Execute(ctx context.Context, req map[string]any) (map[string]any, error) {
	var out map[string]any
	return out, r.h.post(ctx, "/v1/search", req, &out)
}
