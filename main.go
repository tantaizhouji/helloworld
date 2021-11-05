package main

import (
	"fmt"
	"os"
	"strconv"
	"tgpl.com/helloworld/spider"
)

func main() {
	if len(os.Args) > 2 {
		url := os.Args[1]
		times, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "times not int: %s", os.Args[2])
			os.Exit(1)
		}
		doUrl := url
		for i := 0; i < times; i++ {
			fmt.Printf("page %d\n", i + 1)
			next := spider.PrintAndGetNextUrl(doUrl)
			if next == "" {
				break
			}
			doUrl = next
		}
	}
}