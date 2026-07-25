package scrapio

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const defaultBaseURL = "https://api.scrapio.dev"
const defaultTimeout = 30 * time.Second

type httpClient struct {
	base    string
	apiKey  string
	timeout time.Duration
	http    *http.Client
}

func newHTTPClient(base, apiKey string, timeout time.Duration) *httpClient {
	if base == "" {
		base = defaultBaseURL
	}
	if timeout == 0 {
		timeout = defaultTimeout
	}
	return &httpClient{
		base:   base,
		apiKey: apiKey,
		timeout: timeout,
		http:   &http.Client{Timeout: timeout},
	}
}

func (c *httpClient) get(ctx context.Context, path string, query url.Values, out any) error {
	u := c.base + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	return c.do(req, out)
}

func (c *httpClient) post(ctx context.Context, path string, body, out any) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.base+path, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(req, out)
}

func (c *httpClient) do(req *http.Request, out any) error {
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		msg := extractMessage(raw)
		base := ScrapioError{StatusCode: resp.StatusCode, Message: msg}
		switch resp.StatusCode {
		case 401, 403:
			return &AuthError{base}
		case 402:
			return &CreditsExhaustedError{base}
		case 429:
			return &RateLimitError{base}
		default:
			return &base
		}
	}

	if out != nil {
		return json.Unmarshal(raw, out)
	}
	return nil
}

func extractMessage(body []byte) string {
	var m struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}
	if err := json.Unmarshal(body, &m); err == nil {
		if m.Message != "" {
			return m.Message
		}
		if m.Error != "" {
			return m.Error
		}
	}
	s := string(body)
	if len(s) > 200 {
		s = s[:200]
	}
	return fmt.Sprintf("unexpected response: %s", s)
}
