package input

import "strings"

func Clean(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	text = strings.ToLower(text)
	return strings.Fields(text)
}
