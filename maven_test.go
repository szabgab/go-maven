package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestPage(t *testing.T) {
	res, err := parseFile("abc")
	if err == nil {
		t.Errorf("Expected file not found error. Received no error.")
	}
	if !os.IsNotExist(err) {
		t.Errorf("Expected file not found error. Received %v", err)
	}
	if fmt.Sprintf("%T", res) != "main.pageType" {
		t.Errorf("Invalid type: %v", reflect.TypeOf(res))
	}

	path := "pages/a.txt"
	res, err = parseFile(path)
	if err != nil {
		t.Errorf("Error parsing file %v", err)
	}
	if fmt.Sprintf("%T", res) != "main.pageType" {
		t.Errorf("Invalid type: %v", reflect.TypeOf(res))
	}
	if res.title != "Code Maven - for people who code" {
		t.Errorf("Did not read the title correctly from file '%v'. Received: '%v'", path, res.title)
	}
	if res.status != "show" {
		t.Errorf("Did not read the status correctly from file '%v'. Received: '%v'", path, res.status)
	}

}

func TestPagesBadPath(t *testing.T) {
	path := "abc"
	res, err := collectData(path)
	if err == nil {
		t.Errorf("Expected file not found error. Received no error.")
	}
	if !os.IsNotExist(err) {
		t.Errorf("Expected file not found error. Received %v", err)
	}
	if fmt.Sprintf("%T", res) != "[]main.pageType" {
		t.Errorf("Invalid type: %v", reflect.TypeOf(res))
	}
}

func TestPages(t *testing.T) {
	path := "pages"
	res, err := collectData(path)
	if err != nil {
		t.Errorf("Error received: %v", err)
	}
	if fmt.Sprintf("%T", res) != "[]main.pageType" {
		t.Errorf("Invalid type: %v", reflect.TypeOf(res))
	}

}
