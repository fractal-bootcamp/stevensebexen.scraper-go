package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) <= 1 {
		println("Please provide a url.")
		os.Exit(1)
	}

	url := os.Args[1]
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Unable to fetch url %s\n", url)
		os.Exit(1)
	}
	defer res.Body.Close()

	buffer, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading content: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Content length: %d\n", len(buffer))

	doc, err := html.Parse(bytes.NewReader(buffer))
	if err != nil {
		fmt.Printf("Error while parsing: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", doc.FirstChild)
}
