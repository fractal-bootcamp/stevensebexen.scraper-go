package main

import "strings"

// Converts a relative URL to an absolute one.
func rtoa(url string, host string) string {
	if strings.HasPrefix(url, "http") {
		return url
	} else {
		if strings.HasPrefix(url, "/") {
			return httpify(host + url)
		} else {
			return httpify(host + "/" + url)
		}
	}
}
