package utils

import (
	"strings"
)

func UrlFilter(rule string, url string) bool {

	index := strings.Index(rule, "*")
	if index < 0 {
		if rule == url {
			return true
		} else {
			return false
		}
	}

	prefix := rule[:index]

	return strings.HasPrefix(url, prefix)
}
