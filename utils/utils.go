package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-openapi/runtime"
)

// ClientConfig contains configuration passed to the individual API Clients.
type ClientConfig struct {
	// NumRetries contains the number of attempted retries
	NumRetries int
	// RetryTimeout says how long to wait before retrying a request
	RetryTimeout time.Duration
	// RetryStatusCodes contains the list of status codes to retry, use "x" as a wildcard for a single digit (default: [429, 5xx])
	RetryStatusCodes []string
}

func GetTimeout(retryTimeout time.Duration) time.Duration {
	if retryTimeout == 0 {
		return time.Second * 5
	}
	return retryTimeout
}

// MatchRetryCode checks if the status code matches any of the configured retry status codes.
func MatchRetryCode(err error, retryCodes []string) (bool, error) {
	if len(retryCodes) == 0 {
		retryCodes = []string{"429", "5xx"}
	}

	respCode := err.(*runtime.APIError).Code
	respCodeStr := strconv.Itoa(respCode)
	for _, retryCode := range retryCodes {
		if len(retryCode) != 3 {
			return false, fmt.Errorf("invalid retry status code: %s", retryCode)
		}
		matched := true
		for i := range retryCode {
			c := retryCode[i]
			if c == 'x' {
				continue
			}
			if respCodeStr[i] != c {
				matched = false
				break
			}
		}
		if matched {
			return true, nil
		}
	}

	return false, nil
}
