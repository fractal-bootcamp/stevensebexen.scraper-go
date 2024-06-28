package main

import (
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		println("Please provide a url.")
		os.Exit(1)
	}

	url := httpify(os.Args[1])
	doc, host := fetchHtml(url)
	var links []string

	extractLinks(&links, doc, host, 10)
	renderHtml(doc, "render.txt")
	renderLinks(links, "links.txt")

}
