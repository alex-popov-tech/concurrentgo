package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		go func(path string) {
			bytes, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s\n%s\n\n", path, string(bytes))
		}(os.Args[i])
	}
	time.Sleep(time.Second)
}
