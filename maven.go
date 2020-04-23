package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

}

type pageType struct {
	title string
}

func parseFile(path string) (pageType, error) {
	p := pageType{}

	_, err := os.Stat(path)
	if err != nil {
		return p, err
	}

	header := true
	fh, err := os.Open(path)
	if err != nil {
		return p, nil
	}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			header = false
			continue
		}
		if header && strings.HasPrefix(line, "=") {
			parts := strings.SplitN(line[1:], " ", 2)
			if parts[0] == "title" {
				p.title = parts[1]
			}
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		return p, err
	}

	return p, nil
}
