package main

import (
	"golang.org/x/net/html"
)

func extractLinksClosure(links *[]string, host string, linksPerPage int) func(*html.Node) {
	numLinks := 0

	var f func(*html.Node)
	f = func(n *html.Node) {
		if numLinks >= linksPerPage {
			return
		}
		// See if current node has a link.
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, ele := range n.Attr {
				if ele.Key == "href" {
					// Relative links will be printed as absolute ones.
					link := rtoa(ele.Val, host)
					*links = append(*links, link)
					numLinks++
					break
				}
			}
		}

		// Run extractLinks recursively on the first child and its siblings.
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	return func(n *html.Node) {
		f(n)
	}
}

func extractLinks(links *[]string, n *html.Node, host string, linksPerPage int) {
	extractLinksClosure(links, host, linksPerPage)(n)
}
