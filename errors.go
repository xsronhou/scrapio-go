package scrapio

import "fmt"

type ScrapioError struct {
	StatusCode int
	Message    string
}

func (e *ScrapioError) Error() string {
	return fmt.Sprintf("scrapio: HTTP %d: %s", e.StatusCode, e.Message)
}

type AuthError struct{ ScrapioError }
type RateLimitError struct{ ScrapioError }
type CreditsExhaustedError struct{ ScrapioError }
