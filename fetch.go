package scrapio

import "context"

type FetchRequest struct {
	URL       string         `json:"url"`
	RenderJS  *bool          `json:"render_js,omitempty"`
	Device    string         `json:"device,omitempty"`
	Session   *FetchSession  `json:"session,omitempty"`
	Output    []string       `json:"output,omitempty"`
	Extract   map[string]any `json:"extract,omitempty"`
	WaitFor   *FetchWaitFor  `json:"wait_for,omitempty"`
	TimeoutMs *int           `json:"timeout_ms,omitempty"`
}

type FetchSession struct {
	ID string `json:"id"`
}

type FetchWaitFor struct {
	NetworkIdle *bool `json:"network_idle,omitempty"`
}

type FetchResponse struct {
	RequestID   string         `json:"request_id"`
	URL         string         `json:"url"`
	StatusCode  int            `json:"status_code"`
	Outputs     map[string]any `json:"outputs"`
	Diagnostics map[string]any `json:"diagnostics,omitempty"`
}

type FetchResource struct{ h *httpClient }

func (r *FetchResource) Fetch(ctx context.Context, req *FetchRequest) (*FetchResponse, error) {
	var out FetchResponse
	return &out, r.h.post(ctx, "/v1/fetch", req, &out)
}
