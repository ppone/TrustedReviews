package main

import (
	"fmt"
)

var filters = map[string]string{
	"$blank":    "$blank",
	"$bw":       "$bw",
	"$bwin":     "$bwin",
	"$eq":       "$eq",
	"$excludes": "$excludes",
}

func main() {
	fmt.Println("hello")

}
