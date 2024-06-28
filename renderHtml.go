package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func renderHtml(doc *html.Node, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %s", err)
		return
	}

	err = html.Render(f, doc)
	if err != nil {
		fmt.Printf("Error rendering HTML: %s", err)
		return
	}
}
