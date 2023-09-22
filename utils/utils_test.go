package utils_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"

	"github.com/grafana/grafana-openapi-client-go/utils"
)

func TestMatchRetryCode(t *testing.T) {
	tests := []struct {
		name            string
		givenError      error
		givenRetryCodes []string
		expectedRetry   bool
		expectedErr     error
	}{
		{
			name:            "retry with error code 429",
			givenError:      runtime.NewAPIError("", "", http.StatusTooManyRequests),
			givenRetryCodes: []string{},
			expectedRetry:   true,
			expectedErr:     nil,
		},
		{
			name:            "expect retry with error code 500",
			givenError:      runtime.NewAPIError("", "", http.StatusInternalServerError),
			givenRetryCodes: []string{},
			expectedRetry:   true,
			expectedErr:     nil,
		},
		{
			name:            "expect error with invalid retry status code",
			givenError:      runtime.NewAPIError("", "", 1),
			givenRetryCodes: []string{"1"},
			expectedRetry:   false,
			expectedErr:     fmt.Errorf("invalid retry status code"),
		},
		{
			name:            "do not retry or error if given does not match default status to retry",
			givenError:      runtime.NewAPIError("", "", http.StatusAccepted),
			givenRetryCodes: []string{},
			expectedRetry:   false,
			expectedErr:     nil,
		},
		{
			name:            "do not retry or error if given error does not match given status",
			givenError:      runtime.NewAPIError("", "", http.StatusPermanentRedirect),
			givenRetryCodes: []string{"400"},
			expectedRetry:   false,
			expectedErr:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shouldRetry, err := utils.MatchRetryCode(tt.givenError, tt.givenRetryCodes)
			if tt.expectedErr != nil {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.expectedRetry, shouldRetry)
		})
	}
}
