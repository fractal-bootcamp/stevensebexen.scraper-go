package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func renderLinks(links []string, filename string) {
	s := strings.Join(links, "\n")
	w, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %s", err)
	}
	io.WriteString(w, s)
}
