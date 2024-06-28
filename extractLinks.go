package main

import (
	"golang.org/x/net/html"
)

func extractLinksClosure(host string, linksPerPage int) func(*html.Node) []string {
	var links []string

	var f func(*html.Node) []string
	f = func(n *html.Node) []string {
		if len(links) >= linksPerPage {
			return links
		}
		// See if current node has a link.
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, ele := range n.Attr {
				if ele.Key == "href" {
					// Relative links will be printed as absolute ones.
					link := rtoa(ele.Val, host)
					links = append(links, link)
					break
				}
			}
		}

		// Run extractLinks recursively on the first child and its siblings.
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

		return links
	}

	return func(n *html.Node) []string {
		return f(n)
	}
}

func extractLinks(n *html.Node, host string, linksPerPage int) []string {
	return extractLinksClosure(host, linksPerPage)(n)
}
