package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func domainList(filename *string) []string {
	var domains []string

	if *filename != "" {
		file, err := os.Open(*filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		input := bufio.NewScanner(file)

		for input.Scan() {
			if input.Text() == "" {
				break
			}

			domains = append(domains, input.Text())
		}

		file.Close()
	} else {
		domains = flag.Args()
	}

	return domains
}
