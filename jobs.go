package scrapio

import (
	"context"
	"fmt"
	"time"
)

type JobWebhook struct {
	URL string `json:"url"`
}

type CreateJobRequest struct {
	Kind    string         `json:"kind"`
	Input   map[string]any `json:"input"`
	Webhook *JobWebhook    `json:"webhook,omitempty"`
}

type Job struct {
	RequestID  string `json:"request_id"`
	JobID      string `json:"job_id"`
	JobType    string `json:"job_type"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	WebhookURL string `json:"webhook_url,omitempty"`
}

type JobError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type JobResult struct {
	Job
	Result any       `json:"result,omitempty"`
	Error  *JobError `json:"error,omitempty"`
}

var terminalStatuses = map[string]bool{
	"completed": true,
	"partial":   true,
	"failed":    true,
	"cancelled": true,
}

type JobsResource struct{ h *httpClient }

func (r *JobsResource) Create(ctx context.Context, req *CreateJobRequest) (*Job, error) {
	var out Job
	return &out, r.h.post(ctx, "/v1/jobs", req, &out)
}

func (r *JobsResource) Get(ctx context.Context, jobID string) (*Job, error) {
	var out Job
	return &out, r.h.get(ctx, "/v1/jobs/"+jobID, nil, &out)
}

func (r *JobsResource) GetResult(ctx context.Context, jobID string) (*JobResult, error) {
	var out JobResult
	return &out, r.h.get(ctx, "/v1/jobs/"+jobID+"/result", nil, &out)
}

type WaitOptions struct {
	PollInterval time.Duration
	Timeout      time.Duration
}

func (r *JobsResource) WaitForCompletion(ctx context.Context, jobID string, opts *WaitOptions) (*JobResult, error) {
	interval := 2 * time.Second
	deadline := 5 * time.Minute
	if opts != nil {
		if opts.PollInterval > 0 {
			interval = opts.PollInterval
		}
		if opts.Timeout > 0 {
			deadline = opts.Timeout
		}
	}
	ctx, cancel := context.WithTimeout(ctx, deadline)
	defer cancel()

	for {
		job, err := r.Get(ctx, jobID)
		if err != nil {
			return nil, err
		}
		if terminalStatuses[job.Status] {
			return r.GetResult(ctx, jobID)
		}
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("scrapio: job %s did not complete within timeout", jobID)
		case <-time.After(interval):
		}
	}
}
