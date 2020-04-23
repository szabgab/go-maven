package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPage(t *testing.T) {
	res, err := parseFile("abc")
	if err != nil {
		t.Errorf("Error parsing file %v", err)
	}
	//fmt.Println(reflect.TypeOf(res))
	//if reflect.TypeOf(res) != pageType {
	if fmt.Sprintf("%T", res) != "main.pageType" {
		t.Errorf("Invalid type: %v", reflect.TypeOf(res))
	}
}
