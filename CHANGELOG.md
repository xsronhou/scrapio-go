# Changelog

## 1.2.0

- Add `Airbnb`, `Chatgpt`, `Perplexity`, `Gemini`, `Bing`, `Reddit`, `TikTok`, `Target`, and `AppleAppStore` resources on `Client`.
- Add `YouTube.GetChannel`.

## 1.1.1

- Fix `Fetch`, which exposed `Actions`/`Timeout`/`Proxy`/`Country` fields that don't exist on `/v1/fetch` (real fields are `WaitFor`/`TimeoutMs`) -- setting any of them 400'd.
- Fix `YouTube.Search`, `Walmart.Search`, and `Amazon.Search`, which sent `query` instead of the schemas' required `search` -- every call 400'd.
- Fix `Walmart.GetProduct`'s `country` param, which doesn't exist on that schema -- replaced with `zip_code` (latent, since the route is currently disabled server-side).
- Fix `Jobs.Create`, which sent `{JobType, Payload, WebhookURL}` -- the API's `CreateJobRequest` schema is a discriminated union on `{kind, input, webhook: {url}}`. Every call 400'd.

## 1.1.0

- Add `FastSearch`, `Search`, `Map`, `Booking`, and `Agoda` resources on `Client`.
- Fix `v1.0.0` having accidentally shipped the Python SDK's source tree instead of Go code — `go get github.com/xsronhou/scrapio-go@v1.0.0` never actually worked. This is the first version that installs correctly.

## 1.0.0

- Initial release (published with the wrong content — see 1.1.0).
