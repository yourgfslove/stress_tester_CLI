package validation

import (
	"fmt"
	"net/url"
)

var allowedMethods = map[string]struct{}{
	"GET":   {},
	"POST":  {},
	"PUT":   {},
	"PATCH": {},
}

func StressParamsValidate(link, method string, rps, seconds int) error {
	if _, err := url.ParseRequestURI(link); err != nil {
		return fmt.Errorf("failed to parse url")
	}
	if _, ok := allowedMethods[method]; !ok {
		return fmt.Errorf("failed to parse method")
	}
	if rps > 10000 || rps <= 0 {
		return fmt.Errorf("wrong amount of requests per second")
	}
	if seconds <= 0 {
		return fmt.Errorf("wrong amount of seconds")
	}
	return nil
}
