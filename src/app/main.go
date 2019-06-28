package main

import (
	"flag"
	"fmt"
	"os"
	"str"
	"words"
)

func main() {
	fWords := flag.String("count", "", "count words in string, print out most 10 frequent")
	fStr := flag.String("encode", "", "encode string")

	flag.Parse()

	if len(*fWords) > 0 {
		for i, s := range words.Count(*fWords) {
			fmt.Printf("%d - %s\n", i+1, s)
		}
		os.Exit(0)
	}

	if len(*fStr) > 0 {
		fmt.Printf("%s -> %s\n", *fStr, str.Encode(*fStr))
		os.Exit(0)
	}

	flag.Usage()
}
