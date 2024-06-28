package main

import "strings"

func httpify(s string) string {
	if strings.HasPrefix(s, "http") {
		return s
	} else {
		return "http://" + s
	}
}
