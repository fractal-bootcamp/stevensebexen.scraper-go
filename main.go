package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		println("Please provide a url.")
		os.Exit(1)
	}

	url := httpify(os.Args[1])
	doc := fetchHtml(url)
	var links []string

	renderHtml(doc, "render.txt")
	extractLinks(&links, doc, 2, 10)

	s := strings.Join(links, "\n")
	w, err := os.Create("links.txt")
	if err != nil {
		fmt.Printf("Error creating file: %s", err)
	}
	io.WriteString(w, s)
}
