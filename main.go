package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func fetchHtml(url string) *html.Node {
	// HTTP Get
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Unable to fetch url %s\n", url)
		return nil
	}
	defer res.Body.Close()

	var doc *html.Node

	// Use Content-Length if available, otherwise use a buffer to determine length.
	// Then, parse the HTML.
	if res.ContentLength != -1 {
		fmt.Printf("Content length: %d\n", res.ContentLength)
		doc, err = html.Parse(res.Body)
		if err != nil {
			fmt.Printf("Error parsing HTML: %s", err)
		}
	} else {
		buffer, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error reading content: %s\n", err)
			return nil
		}
		fmt.Printf("Content length: %d\n", len(buffer))
		doc, err = html.Parse(bytes.NewBuffer(buffer))
		if err != nil {
			fmt.Printf("Error parsing HTML: %s", err)
		}
	}

	return doc
}

func main() {
	if len(os.Args) <= 1 {
		println("Please provide a url.")
		os.Exit(1)
	}

	url := os.Args[1]
	doc := fetchHtml(url)

	f, err := os.Create("render.txt")
	if err != nil {
		fmt.Printf("Error creating file: %s", err)
		return
	}

	err = html.Render(f, doc)
	if err != nil {
		fmt.Printf("Error rendering HTML: %s", err)
		return
	}
	fmt.Printf("%#v\n", doc.FirstChild)
}
