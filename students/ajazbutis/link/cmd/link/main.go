package main

import (
	"flag"
	"fmt"
	"link"
)

func main() {
	file := flag.String("f", "", "html file to parse, if not provided will return")
	flag.Parse()
	if *file == "" {
		return
	}
	links := link.ExtractLinks(file)
	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}