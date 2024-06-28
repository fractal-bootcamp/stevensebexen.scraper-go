package main

import (
	"golang.org/x/net/html"
)

type DocWithLinks struct {
	doc   *html.Node
	url   string
	links []string
}

func fetchHtmlAndLinks(url string, docsWithLinks *[]DocWithLinks, depth int, linksPerPage int) {
	doc, host := fetchHtml(url)
	if doc == nil {
		return
	}
	links := extractLinks(doc, host, linksPerPage)
	docWithLinks := DocWithLinks{doc, url, links}
	*docsWithLinks = append(*docsWithLinks, docWithLinks)

	if depth <= 0 {
		return
	}

	for _, link := range links {
		fetchHtmlAndLinks(link, docsWithLinks, depth-1, linksPerPage)
	}
}
