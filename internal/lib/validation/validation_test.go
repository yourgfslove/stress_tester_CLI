package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStressParamsValidate(t *testing.T) {
	testcases := []struct {
		name    string
		method  string
		url     string
		rps     int
		seconds int
		err     string
	}{{
		name:    "valid params",
		method:  "GET",
		url:     "https://example.com",
		rps:     50,
		seconds: 10,
	}, {
		name:    "invalid method",
		method:  "ggeeeeetttt",
		url:     "https://example.com",
		rps:     50,
		seconds: 10,
		err:     "failed to parse method",
	}, {
		name:    "invalid url",
		method:  "GET",
		url:     "esadasdxampleasdcom",
		rps:     50,
		seconds: 10,
		err:     "failed to parse url",
	}, {
		name:    "invalid rps",
		method:  "GET",
		url:     "https://example.com",
		rps:     -12,
		seconds: 10,
		err:     "wrong amount of requests per second",
	}, {
		name:    "invalid seconds",
		method:  "GET",
		url:     "https://example.com",
		rps:     50,
		seconds: -10,
		err:     "wrong amount of seconds",
	},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := StressParamsValidate(tc.url, tc.method, tc.rps, tc.seconds)
			if tc.err != "" {
				assert.EqualError(t, err, tc.err, "should err, but another")
			} else {
				assert.NoError(t, err, "expecned no err, but got one")
			}
		})
	}
}
