package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var basePath = "./"
	var seperator = " "
	var fileNum = 0
	var totalFileNum = 0

	err := filepath.Walk(
		basePath,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				totalFileNum++
			}

			return nil
		})

	fmt.Printf("%v Files Discovered\n", totalFileNum)
	if err != nil {
		log.Fatal(err)
	}
	err = filepath.Walk(
		basePath,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				fmt.Println(err)
				return nil
			}

			if !info.IsDir() {
				bytesRead, err := os.ReadFile(path)
				rel, err := filepath.Rel(basePath, path)

				if err != nil {
					fmt.Println("Error reading file")
				}

				convertedPath := filepath.ToSlash(rel)
				finalPath := strings.ReplaceAll(convertedPath, "/", seperator)
				finalPath = filepath.Join(basePath, finalPath)
				err = os.WriteFile(finalPath, bytesRead, 0755)
				fileNum++
				fmt.Printf("Wrote File %s (FileNo. %v / %v)\n", finalPath, fileNum, totalFileNum)
			}

			return nil
		})

	if err != nil {
		log.Fatal(err)
	}

}
