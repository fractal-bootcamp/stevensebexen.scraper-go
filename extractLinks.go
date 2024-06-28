package main

import (
	"golang.org/x/net/html"
)

func extractLinks(links *[]string, n *html.Node, depth int, maxLinks int) bool {
	hasLink := false
	numLinks := 0

	// See if current node has a link
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, ele := range n.Attr {
			if ele.Key == "href" {
				link := httpify(ele.Val)
				*links = append(*links, link)
				hasLink = true
				break
			}
		}
	}

	// Run extractLinks recursively on the first child and its siblings
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if extractLinks(links, c, depth-1, maxLinks) {
			numLinks++
			if numLinks >= maxLinks {
				break
			}
		}
	}

	return hasLink
}
