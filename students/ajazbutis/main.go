package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aJazbutis/link/students/ajazbutis/link"
)

func main() {
	fileFlag := flag.String("f", "", "html file to parse, if not provided will return")
	flag.Parse()
	if *fileFlag == "" {
		return
	}
	file, err := os.Open(*fileFlag)
	if err != nil {
		log.Fatal(err)
	}
	links := link.ExtractLinks(io.Reader(file))
	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}
