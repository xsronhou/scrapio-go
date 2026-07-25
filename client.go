// Package scrapio is the official Go SDK for Scrapio (https://scrapio.dev).
package scrapio

import "time"

// Client is the main entry point for the Scrapio API.
type Client struct {
	Fetch      *FetchResource
	Google     *GoogleResource
	Amazon     *AmazonResource
	Walmart    *WalmartResource
	YouTube    *YouTubeResource
	Jobs       *JobsResource
	Crawl      *CrawlResource
	Interact   *InteractResource
	FastSearch *FastSearchResource
	Search     *SearchResource
	Map        *MapResource
	Booking    *BookingResource
	Agoda      *AgodaResource
}

// ClientOption configures the Client.
type ClientOption func(*clientConfig)

type clientConfig struct {
	baseURL string
	timeout time.Duration
}

// WithBaseURL overrides the default API base URL.
func WithBaseURL(u string) ClientOption {
	return func(c *clientConfig) { c.baseURL = u }
}

// WithTimeout sets the per-request HTTP timeout (default 30s).
func WithTimeout(d time.Duration) ClientOption {
	return func(c *clientConfig) { c.timeout = d }
}

// NewClient creates a new Scrapio client authenticated with apiKey.
func NewClient(apiKey string, opts ...ClientOption) *Client {
	cfg := &clientConfig{}
	for _, o := range opts {
		o(cfg)
	}
	h := newHTTPClient(cfg.baseURL, apiKey, cfg.timeout)
	return &Client{
		Fetch:      &FetchResource{h},
		Google:     &GoogleResource{h},
		Amazon:     &AmazonResource{h},
		Walmart:    &WalmartResource{h},
		YouTube:    &YouTubeResource{h},
		Jobs:       &JobsResource{h},
		Crawl:      &CrawlResource{h},
		Interact:   &InteractResource{h},
		FastSearch: &FastSearchResource{h},
		Search:     &SearchResource{h},
		Map:        &MapResource{h},
		Booking:    &BookingResource{h},
		Agoda:      &AgodaResource{h},
	}
}
