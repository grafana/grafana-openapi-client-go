package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/grafana/grafana-openapi-client-go/client"
)

func retries(cfg *client.Config) error {
	// retry logic
	for n := 0; n <= cfg.NumRetries; n++ {
		req, err := c.newRequest(method, requestPath, query, bytes.NewReader(body))
		if err != nil {
			return err
		}

		// Wait a bit if that's not the first request
		if n != 0 {
			if cfg.RetryTimeout == 0 {
				cfg.RetryTimeout = time.Second * 5
			}
			time.Sleep(cfg.RetryTimeout)
		}

		resp, err = c.client.Do(req)

		// If err is not nil, retry again
		// That's either caused by client policy, or failure to speak HTTP (such as network connectivity problem). A
		// non-2xx status code doesn't cause an error.
		if err != nil {
			continue
		}

		defer resp.Body.Close()

		// read the body (even on non-successful HTTP status codes), as that's what the unit tests expect
		bodyContents, err = ioutil.ReadAll(resp.Body)

		// if there was an error reading the body, try again
		if err != nil {
			continue
		}

		shouldRetry, err := matchRetryCode(resp.StatusCode, cfg.RetryStatusCodes)
		if err != nil {
			return err
		}
		if !shouldRetry {
			break
		}
	}
	if err != nil {
		return err
	}

	return nil
}

// matchRetryCode checks if the status code matches any of the configured retry status codes.
func matchRetryCode(gottenCode int, retryCodes []string) (bool, error) {
	retryStatusCodes := retryCodes
	if len(retryStatusCodes) == 0 {
		retryStatusCodes = []string{"429", "5xx"}
	}

	gottenCodeStr := strconv.Itoa(gottenCode)
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
			if gottenCodeStr[i] != c {
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
