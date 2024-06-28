package main

import (
	"fmt"
	"os"
)

func mkOutDir() {
	_, err := os.Stat("out")

	if os.IsNotExist(err) {
		err := os.Mkdir("out", 0755)
		if err != nil {
			fmt.Printf("Unable to create output directory: %s\n", err)
		}
	}
}
