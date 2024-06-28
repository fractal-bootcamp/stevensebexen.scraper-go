package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseArgs(args []string) (string, int, int, bool) {
	url, depth, linksPerPage := "", DEPTH, LINKS_PER_PAGE
	var err error
	for i := 1; i < len(args); i++ {

		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			argKey := arg[2:]
			i++
			argValue := args[i]

			switch argKey {

			case "depth":
				depth, err = strconv.Atoi(argValue)
				if err != nil {
					fmt.Printf("Invalid value for depth: '%s'. Depth must be an integer.\n", err)
					return url, depth, linksPerPage, false
				}

			case "links":
				linksPerPage, err = strconv.Atoi(argValue)
				if err != nil {
					fmt.Printf("Invalid value for links per page: '%s'. Links per page must be an integer.\n", err)
					return url, depth, linksPerPage, false
				}

			default:
				fmt.Printf("Unrecognized parameter ")

			}
		} else {
			if url != "" {
				fmt.Printf("There must be only one unnamed param for url. Example: scraper google.com --depth 0 --links 5\n")
				return url, depth, linksPerPage, false
			} else {
				url = httpify(arg)
			}
		}
	}

	if url == "" {
		fmt.Println("url is required. Example: scraper google.com --depth 0 --links 5.")
		return url, depth, linksPerPage, false
	}

	return url, depth, linksPerPage, true
}
