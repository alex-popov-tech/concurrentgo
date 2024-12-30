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
			go traverseDir(dirpath+"/"+entry.Name(), searchQuery)
		} else {
			searchInFile(dirpath+"/"+entry.Name(), searchQuery)
		}
	}
	time.Sleep(time.Second)
}

func traverseDir(dirpath string, searchQuery string) {
	entries, err := os.ReadDir(dirpath)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			go traverseDir(dirpath+"/"+entry.Name(), searchQuery)
		} else {
			searchInFile(dirpath+"/"+entry.Name(), searchQuery)
		}
	}
}

func searchInFile(filepath string, searchQuery string) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	if strings.Contains(string(bytes), searchQuery) {
		fmt.Printf("%s - %s\n", filepath, "found match!")
	}
}
