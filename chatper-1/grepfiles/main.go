package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	searchQuery := os.Args[1]
	for i := 2; i < len(os.Args); i++ {
		go func(path string) {
			bytes, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			if strings.Contains(string(bytes), searchQuery) {
				fmt.Printf("%s - %s\n", path, "found match!")
			}
		}(os.Args[i])
	}
	time.Sleep(time.Second)
}
