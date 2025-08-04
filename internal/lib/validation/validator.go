package validation

import (
	"fmt"
	"net/url"
)

const (
	methodGet  = "GET"
	methodPost = "POST"
)

func StressParamsValidate(link string, method string) error {
	if _, err := url.ParseRequestURI(link); err != nil {
		return err
	}
	if method != methodGet && method != methodPost {
		return fmt.Errorf("failed to parse method(GET | POST)")
	}
	return nil
}
