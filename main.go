package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const url_to_scrape string = "http://example.org"

func extract_anchors(n *html.Node, traverse_depth uint, host string, path string) {
	if traverse_depth <= 0 {
		return
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, v := range n.Attr {
			if v.Key == "href" {
				url := v.Val
				if !strings.HasPrefix(v.Val, "http") {
					url = host + v.Val
				}
				fmt.Printf("%s\n", url)
				scrape_page(url, traverse_depth-1)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extract_anchors(c, traverse_depth, host, path)
	}
}

func scrape_page(url string, traverse_depth uint) {
	full_url := url
	if !strings.HasPrefix(full_url, "http") {
		full_url = "http://" + full_url
	}
	response, err := http.Get(full_url)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	if response.StatusCode == 200 {
		document, err := html.Parse(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		host, path := response.Request.URL.Host, response.Request.URL.Path
		extract_anchors(document, traverse_depth, host, path)
	}
}

func main() {
	scrape_page(url_to_scrape, 2)
}
