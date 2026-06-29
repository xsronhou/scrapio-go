# scrapio

Official Python SDK for [Scrapio](https://scrapio.dev) — fetch, crawl, search, and extract structured data from any URL.

## Install

```bash
pip install scrapio-py
```

Requires Python 3.9 or later.

## Quickstart

```python
from scrapio import ApiClient, FetchRequest

client = ApiClient(api_key="YOUR_API_KEY")

result = client.fetch.fetch(FetchRequest(
    url="https://example.com",
    output=["markdown"],
))

print(result.outputs["markdown"])
```

## Usage

### Fetch a page

```python
result = client.fetch.fetch(FetchRequest(
    url="https://news.ycombinator.com",
    render_js=True,
    output=["markdown"],
))
```

### Google Search

```python
from scrapio import GoogleSearchParams

results = client.google.search(GoogleSearchParams(
    search="best web scraping API 2025",
    country_code="us",
))
print(results.organic_results)
```

### Amazon product

```python
product = client.amazon.get_product("B08N5WRWNW")
print(product.title, product.price)
```

### Walmart search

```python
items = client.walmart.search("headphones")
```

### YouTube transcript

```python
video = client.youtube.get_video("dQw4w9WgXcQ")
```

### Browser automation

```python
result = client.interact.interact({
    "url": "https://example.com",
    "actions": [
        {"type": "click", "selector": "#login"},
        {"type": "type", "selector": "#email", "text": "user@example.com"},
    ],
})
```

### Crawl a site

```python
result = client.crawl.crawl({
    "seeds": ["https://docs.example.com"],
    "max_pages": 50,
})
```

### Async jobs

```python
from scrapio import CreateJobRequest

job = client.jobs.create(CreateJobRequest(
    job_type="fetch",
    payload={"url": "https://example.com", "output": ["markdown"]},
))
result = client.jobs.wait_for_completion(job.job_id, poll_interval=2.0, timeout=120.0)
```

### Async client

```python
import asyncio
from scrapio import AsyncApiClient, FetchRequest

async def main():
    async with AsyncApiClient(api_key="YOUR_API_KEY") as client:
        result = await client.fetch.fetch(FetchRequest(
            url="https://example.com",
            output=["markdown"],
        ))
        print(result.outputs["markdown"])

asyncio.run(main())
```

## Configuration

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `api_key` | `str` | required | Your API key |
| `base_url` | `str` | `https://api.scrapio.dev` | Override for local/staging |
| `timeout` | `float` | `30.0` | Per-request timeout (seconds) |
| `max_retries` | `int` | `3` | Max retries on 429/503 |

## Error handling

```python
from scrapio import (
    ApiClient, FetchRequest,
    AuthError, RateLimitError, CreditsExhaustedError, ApiError,
)

try:
    client.fetch.fetch(FetchRequest(url="https://example.com"))
except AuthError:
    print("Invalid API key")
except CreditsExhaustedError:
    print("No credits remaining")
except RateLimitError:
    print("Rate limited — back off and retry")
except ApiError as e:
    print(f"API error {e.status_code}: {e}")
```

## Links

- [Documentation](https://scrapio.dev/docs)
- [API Reference](https://scrapio.dev/docs/api-reference/fetch)
- [Dashboard](https://app.scrapio.dev)
- [Get an API key](https://scrapio.dev#pricing)

## License

MIT
