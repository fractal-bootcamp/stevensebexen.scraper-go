package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func getContentLength(r io.ReadCloser) (int, error) {
	c, err := io.ReadAll(r)
	if err != nil {
		return -1, err
	}
	return len(c), nil
}

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

	contentLength, err := getContentLength(res.Body)
	if contentLength == -1 {
		fmt.Printf("Error reading content: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Content length: %d\n", contentLength)

	_, err = html.Parse(res.Body)
	if err != nil {
		fmt.Printf("Error while parsing: %s\n", err)
		os.Exit(1)
	}
}
