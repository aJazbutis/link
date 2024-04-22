package main

import (
	"flag"
	"fmt"
	"github.com/aJazbutis/link/students/ajazbutis/link"
	"log"
	"os"
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
	links := link.ExtractLinks(file)
	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}
