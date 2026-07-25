package scrapio

import "context"

type InteractAction map[string]any

type InteractRequest struct {
	URL       string         `json:"url"`
	Device    string         `json:"device,omitempty"`
	Session   *FetchSession  `json:"session,omitempty"`
	Actions   []InteractAction `json:"actions"`
	Output    []string       `json:"output,omitempty"`
	Extract   map[string]any `json:"extract,omitempty"`
	TimeoutMS *int           `json:"timeout_ms,omitempty"`
}

type InteractResponse struct {
	RequestID   string         `json:"request_id"`
	Outputs     map[string]any `json:"outputs"`
	Usage       *struct {
		Credits int `json:"credits"`
	} `json:"usage,omitempty"`
	Diagnostics map[string]any `json:"diagnostics,omitempty"`
}

type InteractResource struct{ h *httpClient }

func (r *InteractResource) Interact(ctx context.Context, req *InteractRequest) (*InteractResponse, error) {
	var out InteractResponse
	return &out, r.h.post(ctx, "/v1/interact", req, &out)
}
