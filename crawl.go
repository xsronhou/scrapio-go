package scrapio

import "context"

type CrawlRequest struct {
	Seeds          []string       `json:"seeds"`
	MaxPages       *int           `json:"max_pages,omitempty"`
	MaxDepth       *int           `json:"max_depth,omitempty"`
	SameDomainOnly *bool          `json:"same_domain_only,omitempty"`
	Output         []string       `json:"output,omitempty"`
	Extract        map[string]any `json:"extract,omitempty"`
	TimeoutMS      *int           `json:"timeout_ms,omitempty"`
}

type CrawlPage struct {
	URL           string         `json:"url"`
	Depth         int            `json:"depth"`
	DiscoveredFrom string        `json:"discovered_from,omitempty"`
	Status        string         `json:"status"`
	Outputs       map[string]any `json:"outputs,omitempty"`
	Error         *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

type CrawlSummary struct {
	PagesDiscovered int `json:"pages_discovered"`
	PagesFetched    int `json:"pages_fetched"`
	PagesSucceeded  int `json:"pages_succeeded"`
	PagesFailed     int `json:"pages_failed"`
	PagesSkipped    int `json:"pages_skipped"`
}

type CrawlResponse struct {
	RequestID string `json:"request_id"`
	Status    string `json:"status"`
	Result    struct {
		Seeds   []string     `json:"seeds"`
		Pages   []CrawlPage  `json:"pages"`
		Summary CrawlSummary `json:"summary"`
	} `json:"result"`
}

type CrawlResource struct{ h *httpClient }

func (r *CrawlResource) Crawl(ctx context.Context, req *CrawlRequest) (*CrawlResponse, error) {
	var out CrawlResponse
	return &out, r.h.post(ctx, "/v1/crawl", req, &out)
}
