package main

import (
	"golang.org/x/net/html"
)

func cleanHtml(n *html.Node) {
	var toDelete []*html.Node

	// Mark children for deletion.
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if shouldDeleteNode(c) {
			toDelete = append(toDelete, c)
		}
	}

	// Delete marked children.
	for _, c := range toDelete {
		n.RemoveChild(c)
	}

	// Recursively filter remaining childrens' children.
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cleanHtml(c)
	}
}
