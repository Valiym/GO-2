package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		path = flag.String("path", ".", "dir for search")
		file = flag.String("file", "", "file name for search")
		d    = flag.Bool("d", false, "Delete duplicated files?")
	)
	flag.Parse()

	duplicateList, err := FindDuplicate(*path, *file)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	duplicateCount := len(duplicateList)

	if duplicateCount > 0 {
		fmt.Printf("Found duplicates: %d\n", duplicateCount)
		for i, duplicateName := range duplicateList {
			fmt.Printf("%d. %s\n", i+1, duplicateName)
		}
	} else {
		fmt.Printf("No copy of %q in path: %q\n", *file, *path)
		return
	}

	if *d {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("\nAre you sure you want to delete the duplicates? ('Y' to confirm):\t")
		scanner.Scan()
		userAnswer := scanner.Text()
		if strings.ToLower(userAnswer) == "y" {
			for _, duplicateName := range duplicateList {
				err := os.Remove(duplicateName)
				if err != nil {
					log.Fatalf("Error: %v", err)
				}
				fmt.Println("- deleted:", duplicateName)
			}
		}
	}
}
