package utils

import "strings"

func UrlToFilename(url string) string {
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")
	url = strings.ReplaceAll(url, "/", "_")
	return url
}
