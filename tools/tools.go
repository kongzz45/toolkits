package tools

import (
	"fmt"
	"strings"
)

// URLWithBackSlash 是否将URL带上最后的反斜杠
func URLWithBackSlash(url string, slash bool) string {
	if slash {
		if strings.HasSuffix(url, "/") {
			return url
		}
		return fmt.Sprintf("%s/", url)
	}
	if !strings.HasSuffix(url, "/") {
		return url
	}
	return strings.TrimRight(url, "/")
}
