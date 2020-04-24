package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

}

type pageType struct {
	title  string
	status string
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
			if parts[0] == "status" {
				p.status = parts[1]
			}
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		return p, err
	}

	return p, nil
}

func collectData(path string) ([]pageType, error) {
	pages := []pageType{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return pages, err
	}
	for _, file := range files {
		log.Println(file.Name())
		relPath := filepath.Join(path, file.Name())
		res, err := parseFile(relPath)
		if err != nil {
			return pages, err
		}
		pages = append(pages, res)
	}
	return pages, nil
}
