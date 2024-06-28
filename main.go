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

	// Use Content-Length if available, otherwise use a buffer to determine length
	r := res.Body
	if res.ContentLength != -1 {
		fmt.Printf("Content length: %d\n", res.ContentLength)
	} else {
		buffer, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error reading content: %s\n", err)
			return nil
		}
		r = io.NopCloser(bytes.NewReader(buffer))
		fmt.Printf("Content length: %d\n", len(buffer))
	}

	// Parse body or buffer
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Printf("Error while parsing: %s\n", err)
		return nil
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

	fmt.Printf("%#v\n", doc.FirstChild)
}
