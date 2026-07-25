# scrapio-go

Official Go SDK for [Scrapio](https://scrapio.dev) — fetch, crawl, search, and extract structured data from any URL.

## Install

```bash
go get github.com/xsronhou/scrapio-go
```

Requires Go 1.21 or later.

## Quickstart

```go
package main

import (
    "context"
    "fmt"

    scrapio "github.com/xsronhou/scrapio-go"
)

func main() {
    client := scrapio.NewClient(os.Getenv("SCRAPIO_API_KEY"))

    result, err := client.Fetch.Fetch(context.Background(), &scrapio.FetchRequest{
        URL:    "https://example.com",
        Output: []string{"markdown"},
    })
    if err != nil {
        panic(err)
    }
    fmt.Println(result.Outputs["markdown"])
}
```

## Usage

### Fetch a page

```go
result, err := client.Fetch.Fetch(ctx, &scrapio.FetchRequest{
    URL:    "https://news.ycombinator.com",
    Output: []string{"markdown"},
})
```

### Google Search

```go
results, err := client.Google.Search(ctx, &scrapio.GoogleSearchParams{
    Search:      "best web scraping API 2025",
    CountryCode: "us",
})
fmt.Println(results.Results)
```

### Amazon product

```go
product, err := client.Amazon.GetProduct(ctx, "B08N5WRWNW")
fmt.Println(product.Title, product.Price)
```

### Walmart search

```go
items, err := client.Walmart.Search(ctx, "headphones")
```

### YouTube video

```go
video, err := client.YouTube.GetVideo(ctx, "dQw4w9WgXcQ")
```

### Browser automation

```go
result, err := client.Interact.Interact(ctx, &scrapio.InteractRequest{
    URL: "https://example.com",
    Actions: []scrapio.InteractAction{
        {"type": "click", "selector": "#login"},
        {"type": "type", "selector": "#email", "value": "user@example.com"},
    },
})
```

### Crawl a site

```go
result, err := client.Crawl.Crawl(ctx, &scrapio.CrawlRequest{
    Seeds:    []string{"https://docs.example.com"},
    MaxPages: func(i int) *int { return &i }(50),
    Output:   []string{"markdown"},
})
fmt.Println(result.Result.Summary.PagesSucceeded)
```

### Async jobs

```go
job, err := client.Jobs.Create(ctx, &scrapio.CreateJobRequest{
    JobType: "fetch",
    Payload: map[string]any{
        "url":    "https://example.com",
        "output": []string{"markdown"},
    },
})

result, err := client.Jobs.WaitForCompletion(ctx, job.JobID, &scrapio.WaitOptions{
    PollInterval: 2 * time.Second,
    Timeout:      2 * time.Minute,
})
```

## Configuration

```go
client := scrapio.NewClient(
    "YOUR_API_KEY",
    scrapio.WithBaseURL("https://api.scrapio.dev"),  // optional override
    scrapio.WithTimeout(30 * time.Second),           // optional, default 30s
)
```

## Error handling

```go
result, err := client.Fetch.Fetch(ctx, &scrapio.FetchRequest{URL: "https://example.com"})
if err != nil {
    switch e := err.(type) {
    case *scrapio.AuthError:
        fmt.Println("Invalid API key")
    case *scrapio.CreditsExhaustedError:
        fmt.Println("No credits remaining")
    case *scrapio.RateLimitError:
        fmt.Println("Rate limited — back off and retry")
    case *scrapio.ScrapioError:
        fmt.Printf("API error %d: %s\n", e.StatusCode, e.Message)
    default:
        fmt.Println("Network error:", err)
    }
}
```

## Links

- [Documentation](https://scrapio.dev/docs)
- [API Reference](https://scrapio.dev/docs/api-reference/fetch)
- [Dashboard](https://app.scrapio.dev)
- [Get an API key](https://scrapio.dev#pricing)

## License

MIT
