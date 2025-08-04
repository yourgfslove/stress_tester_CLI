package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStressParamsValidate(t *testing.T) {
	testcases := []struct {
		name   string
		method string
		url    string
		err    string
	}{{
		name:   "valid params",
		method: "GET",
		url:    "https://example.com",
	}, {
		name:   "invalid method",
		method: "ggeeeeetttt",
		url:    "https://example.com",
		err:    "failed to parse method(GET | POST)",
	}, {
		name:   "invalid url",
		method: "GET",
		url:    "esadasdxampleasdcom",
		err:    "failed to parse url",
	},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := StressParamsValidate(tc.url, tc.method)
			if tc.err != "" {
				assert.EqualError(t, err, tc.err, "should err, but another")
			} else {
				assert.NoError(t, err, "expecned no err, but got one")
			}
		})
	}
}
