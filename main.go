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

	var docs []DocWithLinks
	fetchHtmlAndLinks(url, &docs, depth, linksPerPage)

	mkOutDir()
	var links []string
	for i, doc := range docs {
		renderHtml(doc.doc, doc.url, path.Join("out", fmt.Sprintf("%d.html", i)))
		links = append(links, doc.links...)
	}
	renderLinks(links, path.Join("out", "links.txt"))
}
