package main

import (
	"golang.org/x/net/html"
)

func shouldDeleteNode(n *html.Node) bool {
	tagsToFilter := [...]string{"nav", "script"}
	classesToFilter := [...]string{"vector-header", "info-box"}
	idsToFilter := [...]string{"p-lang-btn"}

	if n.Type != html.ElementNode {
		return false
	}

	// Check tag
	for _, tag := range tagsToFilter {
		if n.Data == tag {
			return true
		}
	}

	// Check attributes for class and id
	for _, attr := range n.Attr {
		if attr.Key == "class" {
			for _, class := range classesToFilter {
				if attr.Val == class {
					return true
				}
			}
		} else if attr.Key == "id" {
			for _, id := range idsToFilter {
				if attr.Val == id {
					return true
				}
			}
		}
	}

	// Check for stylesheet
	if n.Data == "link" {
		for _, attr := range n.Attr {
			if attr.Key == "rel" && attr.Val == "stylesheet" {
				return true
			}
		}
	}

	return false
}
