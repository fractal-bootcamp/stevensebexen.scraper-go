package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

// Returns the root node of document and the hostname.
func fetchHtml(url string) (*html.Node, string) {
	// HTTP Get
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Unable to fetch url %s\n", url)
		return nil, ""
	}
	defer res.Body.Close()

	var doc *html.Node

	// Use Content-Length if available, otherwise use a buffer to determine length.
	// Then, parse the HTML.
	if res.ContentLength != -1 {
		fmt.Printf("%s: %d bytes.\n", url, res.ContentLength)
		doc, err = html.Parse(res.Body)
		if err != nil {
			fmt.Printf("Error parsing HTML: %s", err)
		}
	} else {
		buffer, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error reading content: %s\n", err)
			return nil, ""
		}
		fmt.Printf("%s: %d bytes.\n", url, len(buffer))
		doc, err = html.Parse(bytes.NewBuffer(buffer))
		if err != nil {
			fmt.Printf("Error parsing HTML: %s", err)
			return nil, ""
		}
	}

	return doc, res.Request.URL.Host
}
