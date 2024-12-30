package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	searchQuery := os.Args[1]
	dirpath := os.Args[2]
	entries, err := os.ReadDir(dirpath)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		go func(entry os.DirEntry) {
			filename := entry.Name()
			filepath := dirpath + "/" + filename
			bytes, err := os.ReadFile(filepath)
			if err != nil {
				panic(err)
			}
			if strings.Contains(string(bytes), searchQuery) {
				fmt.Printf("%s - %s\n", filepath, "found match!")
			}
		}(entry)
	}
	time.Sleep(time.Second)
}
