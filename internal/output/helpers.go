package output

import (
	"fmt"
	"strings"
	"time"
)

// urlprety deleting prefixes and slashes to make url able to save in file name
func urlprety(URL string) string {
	URL = strings.TrimPrefix(URL, "http://")
	URL = strings.TrimPrefix(URL, "https://")
	URL = strings.ReplaceAll(URL, "/", ".")
	return URL
}

// generating a filename with current time and url
func namegenerator(path, url, outtype string) string {
	preatyURL := urlprety(url)
	return fmt.Sprintf("%s/%v-\"%v\".%v", path, time.Now().Format(timeFormat), preatyURL, outtype)
}
