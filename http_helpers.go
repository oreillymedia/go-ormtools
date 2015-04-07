package ormtools

import (
	"net/http"
	"regexp"
	"strconv"
)

// Function that parses the max-age int from a *http.Response.
// If not found, falls back to fallback.
func MaxAgeHeader(resp *http.Response, fallback int) int {

	r, _ := regexp.Compile(`max-age=(\d+)`)
	cacheControl := resp.Header.Get("Cache-Control")
	submatch := r.FindStringSubmatch(cacheControl)
	maxAge := fallback

	if len(submatch) > 1 {
		convert, err := strconv.Atoi(submatch[1])
		if err == nil {
			maxAge = convert
		}
	}

	return maxAge
}
