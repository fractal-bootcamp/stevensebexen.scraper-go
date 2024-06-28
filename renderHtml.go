package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func renderHtml(doc *html.Node, url string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return
	}

	io.WriteString(f, fmt.Sprintf("<!-- %s -->", url))
	err = html.Render(f, doc)
	if err != nil {
		fmt.Printf("Error rendering HTML: %s\n", err)
		return
	}
}
