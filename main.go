package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	url, depth, linksPerPage, ok := parseArgs(os.Args)
	if !ok {
		os.Exit(1)
	}

	// Recursively fetch pages and their links.
	var docs []DocWithLinks
	fetchHtmlAndLinks(url, &docs, depth, linksPerPage)

	// Filter out unwanted nodes.
	for _, doc := range docs {
		cleanHtml(doc.node)
	}

	// Render out each doc and a list of links containing only unique links.
	mkOutDir()
	linksMap := make(map[string]bool)
	for i, doc := range docs {
		renderHtml(doc.node, doc.url, path.Join("out", fmt.Sprintf("%d.html", i)))
		for _, link := range doc.links {
			linksMap[link] = true
		}
	}
	links := make([]string, 0, len(linksMap))
	for link := range linksMap {
		links = append(links, link)
	}
	renderLinks(links, path.Join("out", "links.txt"))
}
